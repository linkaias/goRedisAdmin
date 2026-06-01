<template>
  <div class="data-viewer">
    <!-- 元信息面板 -->
    <div class="meta-panel">
      <div class="meta-row">
        <span class="meta-label">Key</span>
        <span class="meta-value meta-key">{{ nowInfo.key }}</span>
      </div>
      <div class="meta-row">
        <span class="meta-label">{{ $t('data.type') }}</span>
        <span class="type-badge" :class="'type-' + nowInfo.type">
          <span class="type-icon">{{ typeIcon }}</span>
          {{ nowInfo.type }}
        </span>
      </div>
      <div class="meta-row">
        <span class="meta-label">{{ $t('data.ttl') }}</span>
        <span class="meta-value" :class="{ 'ttl-expired': nowInfo.expire_at === -1 }">
          {{ nowInfo.expire_at === -1 ? $t('data.noExpire') : nowInfo.expire_at }}
        </span>
      </div>
      <div class="meta-row">
        <span class="meta-label">{{ $t('data.size') }}</span>
        <span class="meta-value">{{ nowInfo.len || 0 }} {{ sizeUnit }}</span>
      </div>
    </div>

    <!-- 数据内容区 -->
    <div class="data-content">
      <!-- string -->
      <div v-if="nowInfo.type === 'string'" class="data-section">
        <div class="section-header">
          <span class="section-title">{{ $t('data.value') }}</span>
          <span class="section-info">{{ dataStr.length }} {{ $t('data.chars') }}</span>
          <button class="btn-copy" @click="copyData(dataStr)" :title="$t('data.copy')">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <rect x="9" y="9" width="13" height="13" rx="2" ry="2"/>
              <path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"/>
            </svg>
            {{ $t('data.copy') }}
          </button>
        </div>
        <div class="code-block" :class="{ 'is-json': isJson }">
          <pre>{{ formatString(dataStr) }}</pre>
        </div>
      </div>

      <!-- hash -->
      <div v-else-if="nowInfo.type === 'hash'" class="data-section">
        <div class="section-header">
          <span class="section-title">{{ $t('data.hashFields') }}</span>
          <span class="section-info">{{ data.length }} {{ $t('data.fields') }}</span>
        </div>
        <el-table :data="data" border size="small" max-height="400" class="data-table">
          <el-table-column type="index" width="60" align="center" :label="$t('data.index')" />
          <el-table-column prop="key" :label="$t('data.field')" min-width="150">
            <template slot-scope="{ row }">
              <span class="cell-field">{{ row.key }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="value" :label="$t('data.value')" min-width="200">
            <template slot-scope="{ row }">
              <span class="cell-value">{{ row.value }}</span>
            </template>
          </el-table-column>
          <el-table-column width="60" align="center">
            <template slot-scope="{ row }">
              <button class="btn-copy-mini" @click="copyData(row.key + ': ' + row.value)" :title="$t('data.copy')">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <rect x="9" y="9" width="13" height="13" rx="2" ry="2"/>
                  <path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"/>
                </svg>
              </button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <!-- list -->
      <div v-else-if="nowInfo.type === 'list'" class="data-section">
        <div class="section-header">
          <span class="section-title">{{ $t('data.listMembers') }}</span>
          <span class="section-info">{{ data.length }} {{ $t('data.items') }}</span>
        </div>
        <el-table :data="data" border size="small" max-height="400" class="data-table">
          <el-table-column type="index" width="80" align="center" :label="$t('data.index')" />
          <el-table-column :label="$t('data.value')" min-width="300">
            <template slot-scope="{ row }">
              <span class="cell-value">{{ row }}</span>
            </template>
          </el-table-column>
          <el-table-column width="60" align="center">
            <template slot-scope="{ row }">
              <button class="btn-copy-mini" @click="copyData(row)" :title="$t('data.copy')">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <rect x="9" y="9" width="13" height="13" rx="2" ry="2"/>
                  <path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"/>
                </svg>
              </button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <!-- set -->
      <div v-else-if="nowInfo.type === 'set'" class="data-section">
        <div class="section-header">
          <span class="section-title">{{ $t('data.setMembers') }}</span>
          <span class="section-info">{{ data.length }} {{ $t('data.members') }}</span>
        </div>
        <el-table :data="data" border size="small" max-height="400" class="data-table">
          <el-table-column type="index" width="80" align="center" :label="$t('data.index')" />
          <el-table-column :label="$t('data.member')" min-width="300">
            <template slot-scope="{ row }">
              <span class="cell-value">{{ row }}</span>
            </template>
          </el-table-column>
          <el-table-column width="60" align="center">
            <template slot-scope="{ row }">
              <button class="btn-copy-mini" @click="copyData(row)" :title="$t('data.copy')">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <rect x="9" y="9" width="13" height="13" rx="2" ry="2"/>
                  <path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"/>
                </svg>
              </button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <!-- zset -->
      <div v-else-if="nowInfo.type === 'zset'" class="data-section">
        <div class="section-header">
          <span class="section-title">{{ $t('data.zsetMembers') }}</span>
          <span class="section-info">{{ data.length }} {{ $t('data.members') }}</span>
        </div>
        <el-table :data="data" border size="small" max-height="400" class="data-table">
          <el-table-column type="index" width="60" align="center" :label="$t('data.index')" />
          <el-table-column prop="member" :label="$t('data.member')" min-width="200">
            <template slot-scope="{ row }">
              <span class="cell-value">{{ row.member }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="score" :label="$t('data.score')" width="120" align="right">
            <template slot-scope="{ row }">
              <span class="cell-score">{{ row.score }}</span>
            </template>
          </el-table-column>
          <el-table-column width="60" align="center">
            <template slot-scope="{ row }">
              <button class="btn-copy-mini" @click="copyData(row.member)" :title="$t('data.copy')">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <rect x="9" y="9" width="13" height="13" rx="2" ry="2"/>
                  <path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"/>
                </svg>
              </button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <!-- stream -->
      <div v-else-if="nowInfo.type === 'stream'" class="data-section">
        <div class="section-header">
          <span class="section-title">{{ $t('data.streamEntries') }}</span>
          <span class="section-info">{{ streamData.length }} {{ $t('data.entries') }}</span>
        </div>
        <el-table :data="streamData" border size="small" max-height="400" class="data-table">
          <el-table-column type="index" width="60" align="center" :label="$t('data.index')" />
          <el-table-column prop="id" label="ID" width="200">
            <template slot-scope="{ row }">
              <span class="cell-id">{{ row.id }}</span>
            </template>
          </el-table-column>
          <el-table-column :label="$t('data.fields')" min-width="300">
            <template slot-scope="{ row }">
              <div class="stream-fields">
                <div v-for="(val, key, idx) in parseFields(row.fields)" :key="idx" class="stream-field-row">
                  <span class="stream-field-key">{{ key }}:</span>
                  <span class="stream-field-val">{{ val }}</span>
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column width="60" align="center">
            <template slot-scope="{ row }">
              <button class="btn-copy-mini" @click="copyData(JSON.stringify(row, null, 2))" :title="$t('data.copy')">
                <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <rect x="9" y="9" width="13" height="13" rx="2" ry="2"/>
                  <path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"/>
                </svg>
              </button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <!-- unsupported -->
      <div v-else class="data-section">
        <div class="unsupported-box">
          <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="10"/>
            <line x1="12" y1="8" x2="12" y2="12"/>
            <line x1="12" y1="16" x2="12.01" y2="16"/>
          </svg>
          <p>{{ $t('data.unsupportedWithType', { type: nowInfo.type }) }}</p>
        </div>
      </div>
    </div>

    <!-- 底部操作栏 -->
    <div class="data-footer">
      <div class="footer-left">
        <button class="btn-copy-all" @click="copyAll" :title="$t('data.copyAll')">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="9" y="9" width="13" height="13" rx="2" ry="2"/>
            <path d="M5 15H4a2 2 0 01-2-2V4a2 2 0 012-2h9a2 2 0 012 2v1"/>
          </svg>
          {{ $t('data.copyAll') }}
        </button>
      </div>
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
  name: "dataPage",
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
  computed: {
    typeIcon() {
      const icons = {
        string: 'T',
        hash: '#',
        list: '[]',
        set: '{}',
        zset: '{}',
        stream: '>>'
      }
      return icons[this.nowInfo.type] || '?'
    },
    sizeUnit() {
      const type = this.nowInfo.type
      if (type === 'string') return this.$t('data.chars')
      if (type === 'hash') return this.$t('data.fields')
      if (type === 'list' || type === 'set' || type === 'zset') return this.$t('data.members')
      if (type === 'stream') return this.$t('data.entries')
      return ''
    },
    isJson() {
      if (!this.dataStr) return false
      try {
        JSON.parse(this.dataStr)
        return true
      } catch {
        return false
      }
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
    },
    formatString(str) {
      if (!str) return ''
      try {
        const parsed = JSON.parse(str)
        return JSON.stringify(parsed, null, 2)
      } catch {
        return str
      }
    },
    parseFields(fields) {
      if (typeof fields === 'object') return fields
      try {
        return JSON.parse(fields)
      } catch {
        return { raw: fields }
      }
    },
    copyData(text) {
      navigator.clipboard.writeText(text).then(() => {
        this.$message.success(this.$t('data.copied'))
      }).catch(() => {
        this.$message.error(this.$t('data.copyFailed'))
      })
    },
    copyAll() {
      let text = ''
      const type = this.nowInfo.type
      if (type === 'string') {
        text = this.dataStr
      } else if (type === 'hash') {
        text = this.data.map(item => `${item.key}: ${item.value}`).join('\n')
      } else if (type === 'zset') {
        text = this.data.map(item => `${item.score}\t${item.member}`).join('\n')
      } else if (type === 'stream') {
        text = JSON.stringify(this.streamData, null, 2)
      } else {
        text = this.data.join('\n')
      }
      this.copyData(text)
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

/* ── Meta Panel ── */
.meta-panel {
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--border-subtle);
  border-radius: 8px;
  padding: 14px 16px;
  margin-bottom: 20px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
}

.meta-row {
  display: flex;
  align-items: center;
  gap: 10px;
}

.meta-label {
  font-family: var(--font-mono);
  font-size: 11px;
  font-weight: 500;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
  min-width: 40px;
}

.meta-value {
  font-family: var(--font-mono);
  font-size: 13px;
  color: var(--text-secondary);
}

.meta-key {
  color: var(--accent-light);
  font-weight: 500;
  word-break: break-all;
}

.ttl-expired {
  color: var(--text-muted);
  font-style: italic;
}

.type-badge {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 3px 10px;
  border-radius: 5px;
  font-family: var(--font-mono);
  font-size: 12px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.type-icon {
  font-size: 14px;
  opacity: 0.8;
}

.type-string { background: rgba(76, 175, 125, 0.15); color: #4caf7d; }
.type-hash   { background: rgba(220, 107, 47, 0.15); color: #dc6b2f; }
.type-list   { background: rgba(100, 149, 237, 0.15); color: #6495ed; }
.type-set    { background: rgba(187, 134, 252, 0.15); color: #bb86fc; }
.type-zset   { background: rgba(255, 215, 0, 0.15); color: #daa520; }
.type-stream { background: rgba(64, 196, 255, 0.15); color: #40c4ff; }

/* ── Data Content ── */
.data-content {
  margin-bottom: 20px;
}

.data-section {
  animation: section-enter 0.2s ease both;
}

@keyframes section-enter {
  from { opacity: 0; }
}

.section-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 12px;
  padding-bottom: 10px;
  border-bottom: 1px solid var(--border-subtle);
}

.section-title {
  font-family: var(--font-mono);
  font-size: 12px;
  font-weight: 600;
  color: var(--text-primary);
  text-transform: uppercase;
  letter-spacing: 0.8px;
}

.section-info {
  font-family: var(--font-mono);
  font-size: 11px;
  color: var(--text-muted);
  background: rgba(255, 255, 255, 0.05);
  padding: 2px 8px;
  border-radius: 4px;
}

.btn-copy {
  margin-left: auto;
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 4px 10px;
  border: 1px solid var(--border-light);
  border-radius: 5px;
  background: transparent;
  color: var(--text-muted);
  font-family: var(--font-mono);
  font-size: 11px;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-copy:hover {
  border-color: var(--accent);
  color: var(--accent);
  background: var(--accent-glow);
}

.btn-copy-mini {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 28px;
  height: 28px;
  border: none;
  border-radius: 4px;
  background: transparent;
  color: var(--text-muted);
  cursor: pointer;
  transition: all 0.15s;
}

.btn-copy-mini:hover {
  background: rgba(255, 255, 255, 0.08);
  color: var(--accent);
}

/* ── Code Block (string) ── */
.code-block {
  font-family: var(--font-mono);
  font-size: 13px;
  color: var(--accent-light);
  background: rgba(0, 0, 0, 0.3);
  border: 1px solid var(--border-subtle);
  border-radius: 8px;
  padding: 14px 16px;
  max-height: 400px;
  overflow: auto;
}

.code-block pre {
  margin: 0;
  white-space: pre-wrap;
  word-break: break-all;
  line-height: 1.6;
}

.code-block.is-json {
  color: #a8d4a8;
}

/* ── Data Table ── */
.data-table {
  width: 100%;
  border-radius: 8px;
  overflow: hidden;
}

.data-table >>> .el-table__header th {
  background: rgba(255, 255, 255, 0.03);
  font-family: var(--font-mono);
  font-size: 11px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--text-muted);
}

.data-table >>> .el-table__body td {
  border-bottom-color: var(--border-subtle);
}

.cell-field {
  font-family: var(--font-mono);
  font-size: 12px;
  color: var(--accent-light);
  font-weight: 500;
}

.cell-value {
  font-family: var(--font-mono);
  font-size: 12px;
  color: var(--text-secondary);
  word-break: break-all;
}

.cell-score {
  font-family: var(--font-mono);
  font-size: 13px;
  color: #daa520;
  font-weight: 600;
}

.cell-id {
  font-family: var(--font-mono);
  font-size: 11px;
  color: var(--text-muted);
}

/* ── Stream Fields ── */
.stream-fields {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.stream-field-row {
  display: flex;
  gap: 8px;
  font-size: 12px;
}

.stream-field-key {
  font-family: var(--font-mono);
  color: var(--accent-light);
  font-weight: 500;
  min-width: fit-content;
}

.stream-field-val {
  font-family: var(--font-mono);
  color: var(--text-secondary);
  word-break: break-all;
}

/* ── Unsupported ── */
.unsupported-box {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 0;
  color: var(--text-muted);
}

.unsupported-box svg {
  margin-bottom: 12px;
  opacity: 0.4;
}

.unsupported-box p {
  font-size: 14px;
  margin: 0;
}

/* ── Footer ── */
.data-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding-top: 16px;
  border-top: 1px solid var(--border-subtle);
}

.footer-left {
  display: flex;
  gap: 8px;
}

.btn-copy-all {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 7px 14px;
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

.btn-copy-all:hover {
  border-color: var(--accent);
  color: var(--accent);
  background: var(--accent-glow);
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
  background: rgba(255, 255, 255, 0.04);
}
</style>
