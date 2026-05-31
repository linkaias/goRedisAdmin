<template>
  <div class="header-inner">
    <div class="header-brand" @click="$router.push('/')">
      <span class="brand-bracket">[</span>
      <span class="brand-text">redis</span>
      <span class="brand-at">@</span>
      <span class="brand-host">admin</span>
      <span class="brand-bracket">]</span>
    </div>

    <nav class="header-nav">
      <a
        class="nav-item"
        :class="{ 'is-active': activeRoute === '/home' || activeRoute === '/' }"
        @click="go('/')"
      >
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
          <ellipse cx="12" cy="5" rx="9" ry="3"/>
          <path d="M21 12c0 1.66-4 3-9 3s-9-1.34-9-3"/>
          <path d="M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5"/>
        </svg>
        <span>数据库</span>
      </a>
      <a
        class="nav-item"
        :class="{ 'is-active': activeRoute === '/info' }"
        @click="go('/info')"
      >
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10"/>
          <line x1="12" y1="16" x2="12" y2="12"/>
          <line x1="12" y1="8" x2="12.01" y2="8"/>
        </svg>
        <span>Redis 信息</span>
      </a>
    </nav>

    <div class="header-actions">
      <a class="action-link" href="https://github.com/linkaias/goRedisAdmin" target="_blank" title="GitHub">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
          <path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"/>
        </svg>
      </a>
      <div class="action-divider"></div>
      <a class="nav-item nav-logout" @click="logout">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
          <path d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4"/>
          <polyline points="16 17 21 12 16 7"/>
          <line x1="21" y1="12" x2="9" y2="12"/>
        </svg>
        <span>退出</span>
      </a>
    </div>
  </div>
</template>

<script>
import {RemoveToken} from "@/utils/token";

export default {
  name: "HeaderItem",
  computed: {
    activeRoute() {
      return this.$route.path
    }
  },
  methods: {
    go(path) {
      if (this.$route.path !== path) {
        this.$router.push(path)
      }
    },
    logout() {
      RemoveToken()
      location.reload()
    }
  }
}
</script>

<style scoped>
.header-inner {
  display: flex;
  align-items: center;
  height: 60px;
  max-width: 1400px;
  margin: 0 auto;
}

/* ── Brand ── */
.header-brand {
  font-family: var(--font-mono);
  font-size: 17px;
  font-weight: 600;
  cursor: pointer;
  margin-right: 40px;
  flex-shrink: 0;
  user-select: none;
  transition: opacity 0.2s;
}

.header-brand:hover {
  opacity: 0.8;
}

.brand-bracket { color: var(--text-muted); }
.brand-text    { color: var(--text-secondary); }
.brand-at      { color: var(--accent); }
.brand-host    { color: var(--accent-light); }

/* ── Nav ── */
.header-nav {
  display: flex;
  gap: 4px;
  flex: 1;
}

.nav-item {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 8px 14px;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 500;
  color: var(--text-secondary);
  cursor: pointer;
  transition: all 0.2s ease;
  text-decoration: none;
  user-select: none;
}

.nav-item:hover {
  color: var(--text-primary);
  background: rgba(255,255,255,0.04);
}

.nav-item.is-active {
  color: var(--accent-light);
  background: var(--accent-glow);
}

/* ── Actions ── */
.header-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.action-link {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border-radius: 8px;
  color: var(--text-muted);
  transition: all 0.2s;
}

.action-link:hover {
  color: var(--text-primary);
  background: rgba(255,255,255,0.04);
}

.action-divider {
  width: 1px;
  height: 20px;
  background: var(--border-subtle);
}

.nav-logout {
  color: var(--text-muted);
}

.nav-logout:hover {
  color: var(--danger);
  background: rgba(224, 82, 82, 0.08);
}
</style>
