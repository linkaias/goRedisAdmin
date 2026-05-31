<template>
  <div class="data-viewer">
    <div class="data-header">
      <span class="data-key">{{ nowInfo.key }}</span>
      <span class="type-badge" :class="'type-' + nowInfo.type">{{ nowInfo.type }}</span>
    </div>

    <!-- string -->
    <div v-if="nowInfo.type === 'string'" class="data-section">
      <div class="section-label">{{ $t('data.value') }}</div>
      <div class="code-block">{{ dataStr }}</div>
    </div>

    <!-- hash -->
    <div v-else-if="nowInfo.type === 'hash'" class="data-section">
      <div class="section-label">{{ $t('data.hashFields') }} <span class="count-badge">{{ data.length }}</span></div>
      <div class="hash-table">
        <div class="hash-row hash-row-header">
          <span class="hash-col hash-idx">{{ $t('data.index') }}</span>
          <span class="hash-col hash-key">{{ $t('data.key') }}</span>
          <span class="hash-col hash-val">{{ $t('data.value') }}</span>
        </div>
        <div class="hash-row" v-for="(item, idx) in data" :key="idx">
          <span class="hash-col hash-idx">{{ idx + 1 }}</span>
          <span class="hash-col hash-key">{{ item.key }}</span>
          <span class="hash-col hash-val">{{ item.value }}</span>
        </div>
      </div>
    </div>

    <!-- stream -->
    <div v-else-if="nowInfo.type === 'stream'" class="data-section">
      <div class="section-label">{{ $t('data.streamEntries') }} <span class="count-badge">{{ streamData.length }}</span></div>
      <div class="hash-table">
        <div class="hash-row hash-row-header">
          <span class="hash-col stream-id">ID</span>
          <span class="hash-col stream-fields">{{ $t('data.fields') }}</span>
        </div>
        <div class="hash-row" v-for="(item, idx) in streamData" :key="idx">
          <span class="hash-col stream-id">{{ item.id }}</span>
          <span class="hash-col stream-fields">{{ item.fields }}</span>
        </div>
      </div>
    </div>

    <!-- set / list / zset -->
    <div v-else-if="nowInfo.type === 'set' || nowInfo.type === 'list' || nowInfo.type === 'zset'" class="data-section">
      <div class="section-label">{{ $t('data.members', { type: nowInfo.type.toUpperCase() }) }} <span class="count-badge">{{ data.length }}</span></div>
      <div class="tag-list">
        <span class="data-tag" v-for="(item, idx) in data" :key="idx">{{ item }}</span>
      </div>
    </div>

    <!-- unsupported -->
    <div v-else class="data-section">
      <p style="color: var(--text-muted)">{{ $t('data.unsupportedWithType', { type: nowInfo.type }) }}</p>
    </div>

    <div class="data-footer">
      <button class="btn-close" @click="close">
        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <line x1="18" y1="6" x2="6" y2="18"/>
          <line x1="6" y1="6" x2="18" y2="18"/>
        </svg>
        {{ $t('common.close') }}
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: "data",
  data() {
    return {
      dbNum: -1,
      nowInfo: {},
      data: [],
      streamData: [],
      dataStr: "",
      cursor: 0,
      count: 0,
    }
  },
  methods: {
    initData(dbNum, info) {
      this.dbNum = dbNum
      this.nowInfo = info
      this.getData()
    },
    close() {
      this.$emit("closeData")
    },
    async getData() {
      let data = { key: this.nowInfo.key, type: this.nowInfo.type }
      let res = await this.$API.dbApi.reqGetValueByKey(this.dbNum, data)
      if (res.code === 0) {
        let payload = res.data || {}
        this.data = []
        this.streamData = []
        this.dataStr = ""
        if (this.nowInfo.type === "string") {
          this.dataStr = payload || ""
        } else if (this.nowInfo.type === "stream") {
          this.streamData = Array.isArray(payload.data) ? payload.data : []
        } else {
          this.data = Array.isArray(payload.data) ? payload.data : []
        }
        this.cursor = payload.cursor || 0
        this.count = payload.count || 0
      }
    }
  }
}
</script>

<style scoped>
.data-viewer {
  animation: fade-in 0.3s ease both;
}

@keyframes fade-in {
  from { opacity: 0; transform: translateY(8px); }
}

.data-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 20px;
}

.data-key {
  font-family: var(--font-mono);
  font-size: 15px;
  font-weight: 600;
  color: var(--text-primary);
  word-break: break-all;
}

.type-badge {
  display: inline-block;
  padding: 2px 8px;
  border-radius: 4px;
  font-family: var(--font-mono);
  font-size: 11px;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.3px;
  flex-shrink: 0;
}

.type-string { background: rgba(76, 175, 125, 0.12); color: #4caf7d; }
.type-hash   { background: rgba(220, 107, 47, 0.12); color: #dc6b2f; }
.type-list   { background: rgba(100, 149, 237, 0.12); color: #6495ed; }
.type-set    { background: rgba(187, 134, 252, 0.12); color: #bb86fc; }
.type-zset   { background: rgba(255, 215, 0, 0.12); color: #daa520; }
.type-stream { background: rgba(64, 196, 255, 0.12); color: #40c4ff; }

.data-section {
  margin-bottom: 20px;
}

.section-label {
  font-family: var(--font-mono);
  font-size: 11px;
  font-weight: 500;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 1px;
  margin-bottom: 10px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.count-badge {
  font-size: 10px;
  background: rgba(255,255,255,0.06);
  padding: 1px 6px;
  border-radius: 3px;
  color: var(--text-secondary);
}

.code-block {
  font-family: var(--font-mono);
  font-size: 13px;
  color: var(--accent-light);
  background: rgba(0,0,0,0.3);
  border: 1px solid var(--border-subtle);
  border-radius: 8px;
  padding: 14px 16px;
  word-break: break-all;
  line-height: 1.6;
  max-height: 300px;
  overflow-y: auto;
}

/* ── Hash table ── */
.hash-table {
  border: 1px solid var(--border-subtle);
  border-radius: 8px;
  overflow: hidden;
  max-height: 360px;
  overflow-y: auto;
}

.hash-row {
  display: flex;
  border-bottom: 1px solid var(--border-subtle);
}

.hash-row:last-child {
  border-bottom: none;
}

.hash-row-header {
  background: rgba(255,255,255,0.03);
}

.hash-row-header .hash-col {
  font-family: var(--font-mono);
  font-size: 11px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--text-muted);
  font-weight: 500;
}

.hash-col {
  padding: 8px 12px;
  font-size: 13px;
  color: var(--text-primary);
  word-break: break-all;
}

.hash-idx {
  width: 50px;
  flex-shrink: 0;
  text-align: center;
  color: var(--text-muted);
  font-family: var(--font-mono);
  font-size: 11px;
}

.hash-key {
  width: 140px;
  flex-shrink: 0;
  font-family: var(--font-mono);
  font-size: 12px;
  color: var(--accent-light);
  border-right: 1px solid var(--border-subtle);
}

.hash-val {
  flex: 1;
  min-width: 0;
}

.stream-id {
  width: 210px;
  flex-shrink: 0;
  font-family: var(--font-mono);
  font-size: 12px;
  color: var(--accent-light);
  border-right: 1px solid var(--border-subtle);
}

.stream-fields {
  flex: 1;
  min-width: 0;
  word-break: break-all;
}

/* ── Tag list ── */
.tag-list {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  max-height: 360px;
  overflow-y: auto;
}

.data-tag {
  display: inline-block;
  padding: 5px 12px;
  background: rgba(255,255,255,0.04);
  border: 1px solid var(--border-subtle);
  border-radius: 6px;
  font-size: 13px;
  color: var(--text-secondary);
  font-family: var(--font-mono);
  word-break: break-all;
  transition: all 0.15s;
}

.data-tag:hover {
  border-color: rgba(220, 107, 47, 0.3);
  background: var(--accent-glow);
  color: var(--accent-light);
}

/* ── Footer ── */
.data-footer {
  display: flex;
  justify-content: flex-end;
  padding-top: 16px;
  border-top: 1px solid var(--border-subtle);
}

.btn-close {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 7px 16px;
  border: 1px solid var(--border-light);
  border-radius: 7px;
  background: transparent;
  color: var(--text-secondary);
  font-family: var(--font-body);
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-close:hover {
  border-color: var(--text-muted);
  color: var(--text-primary);
  background: rgba(255,255,255,0.04);
}
</style>
