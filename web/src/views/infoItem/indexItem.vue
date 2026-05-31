<template>
  <div class="info-page">
    <div class="info-header">
      <div class="info-title">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10"/>
          <line x1="12" y1="16" x2="12" y2="12"/>
          <line x1="12" y1="8" x2="12.01" y2="8"/>
        </svg>
        <span>{{ $t('info.title') }}</span>
      </div>
      <button class="btn-refresh" @click="getInfo" :class="{ 'is-spinning': isLoading }">
        <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="23 4 23 10 17 10"/>
          <path d="M20.49 15a9 9 0 1 1-2.12-9.36L23 10"/>
        </svg>
        {{ $t('common.refresh') }}
      </button>
    </div>

    <div class="info-body" v-loading="isLoading">
      <div class="terminal-window">
        <div class="terminal-bar">
          <span class="terminal-dot terminal-dot--r"></span>
          <span class="terminal-dot terminal-dot--y"></span>
          <span class="terminal-dot terminal-dot--g"></span>
          <span class="terminal-title">redis-cli INFO</span>
        </div>
        <div class="terminal-content" v-html="formattedInfo"></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "indexItem",
  data() {
    return {
      isLoading: false,
      InfoData: ""
    }
  },
  computed: {
    formattedInfo() {
      if (!this.InfoData) return `<span style="color: var(--text-muted)">${this.$t('info.waitLoading')}</span>`
      return this.InfoData
        .replace(/# ([^\n<]+)/g, '<span class="info-section">$1</span>')
        .replace(/([a-z_]+):/g, '<span class="info-key">$1</span>:')
    }
  },
  mounted() {
    this.getInfo()
  },
  methods: {
    async getInfo() {
      this.isLoading = true
      let res = await this.$API.InfoApi.reqRedisInfo()
      this.InfoData = res.data
      this.isLoading = false
    }
  }
}
</script>

<style scoped>
.info-page {
  max-width: 900px;
  margin: 0 auto;
  animation: page-enter 0.4s cubic-bezier(0.16, 1, 0.3, 1) both;
}

@keyframes page-enter {
  from { opacity: 0; transform: translateY(12px); }
}

.info-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

.info-title {
  display: flex;
  align-items: center;
  gap: 10px;
  font-family: var(--font-mono);
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
}

.info-title svg {
  color: var(--accent);
}

.btn-refresh {
  display: inline-flex;
  align-items: center;
  gap: 6px;
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

.btn-refresh:hover {
  border-color: var(--accent);
  color: var(--accent);
  background: var(--accent-glow);
}

.btn-refresh.is-spinning svg {
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* ── Terminal ── */
.terminal-window {
  background: rgba(0, 0, 0, 0.4);
  border: 1px solid var(--border-subtle);
  border-radius: 12px;
  overflow: hidden;
}

.terminal-bar {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 10px 14px;
  background: rgba(255,255,255,0.03);
  border-bottom: 1px solid var(--border-subtle);
}

.terminal-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}

.terminal-dot--r { background: #ff5f57; }
.terminal-dot--y { background: #febc2e; }
.terminal-dot--g { background: #28c840; }

.terminal-title {
  margin-left: 10px;
  font-family: var(--font-mono);
  font-size: 12px;
  color: var(--text-muted);
}

.terminal-content {
  padding: 20px 24px;
  font-family: var(--font-mono);
  font-size: 12.5px;
  line-height: 1.8;
  color: var(--text-secondary);
  max-height: 70vh;
  overflow-y: auto;
  white-space: pre-wrap;
  word-break: break-all;
}

/* Deep selectors for dynamic HTML content */
.terminal-content >>> .info-section {
  display: block;
  color: var(--accent);
  font-weight: 600;
  font-size: 13px;
  margin: 16px 0 6px;
  padding-bottom: 4px;
  border-bottom: 1px solid rgba(220, 107, 47, 0.15);
}

.terminal-content >>> .info-section:first-child {
  margin-top: 0;
}

.terminal-content >>> .info-key {
  color: var(--accent-light);
}
</style>
