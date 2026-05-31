<template>
  <div class="login-page">
    <!-- 背景装饰 -->
    <div class="bg-grid"></div>
    <div class="bg-glow bg-glow--1"></div>
    <div class="bg-glow bg-glow--2"></div>

    <!-- 登录卡片 -->
    <div class="login-card" :class="{ 'is-loading': loading, 'is-success': success }">
      <!-- 顶部装饰条 -->
      <div class="card-accent"></div>

      <!-- 头部 -->
      <div class="card-header">
        <div class="logo">
          <span class="logo-bracket">[</span>
          <span class="logo-text">redis</span>
          <span class="logo-at">@</span>
          <span class="logo-host">admin</span>
          <span class="logo-bracket">]</span>
          <span class="logo-cursor">_</span>
        </div>
        <p class="subtitle">连接到管理面板</p>
      </div>

      <!-- 表单 -->
      <div class="card-body" v-show="!success">
        <div class="input-group">
          <label class="input-label">用户名</label>
          <div class="input-wrap" :class="{ 'is-focus': focusUser }">
            <span class="input-icon">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
                <circle cx="12" cy="7" r="4"/>
              </svg>
            </span>
            <input
              v-model="user"
              type="text"
              placeholder="输入用户名"
              @focus="focusUser = true"
              @blur="focusUser = false"
              @keyup.enter="$refs.pwdInput.focus()"
            />
          </div>
        </div>

        <div class="input-group">
          <label class="input-label">密码</label>
          <div class="input-wrap" :class="{ 'is-focus': focusPwd }">
            <span class="input-icon">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="11" width="18" height="11" rx="2" ry="2"/>
                <path d="M7 11V7a5 5 0 0 1 10 0v4"/>
              </svg>
            </span>
            <input
              ref="pwdInput"
              v-model="pwd"
              type="password"
              placeholder="输入密码"
              @focus="focusPwd = true"
              @blur="focusPwd = false"
              @keyup.enter="login"
            />
          </div>
        </div>

        <button class="btn-login" @click="login" :disabled="loading">
          <span class="btn-text" v-show="!loading">登 录</span>
          <span class="btn-spinner" v-show="loading">
            <span class="spinner"></span>
            验证中...
          </span>
        </button>
      </div>

      <!-- 成功状态 -->
      <div class="card-success" v-show="success">
        <div class="success-icon">
          <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="20 6 9 17 4 12"/>
          </svg>
        </div>
        <p class="success-text">连接成功</p>
        <p class="success-sub">正在进入管理面板...</p>
      </div>

      <!-- 底部 -->
      <div class="card-footer">
        <span class="footer-dot"></span>
        <span>GoRedisAdmin</span>
      </div>
    </div>
  </div>
</template>

<script>
import {SetToken} from "@/utils/token";

export default {
  name: "login",
  data() {
    return {
      user: "admin",
      pwd: "123456",
      loading: false,
      success: false,
      focusUser: false,
      focusPwd: false,
    }
  },
  methods: {
    async login() {
      if (this.user === "" || this.pwd === "") {
        this.$message.error("用户名或密码不能为空")
        return
      }
      this.loading = true
      try {
        let res = await this.$API.dbApi.reqLogin({user: this.user, pwd: this.pwd})
        if (res.code === 0) {
          SetToken(res.data)
          this.loading = false
          this.success = true
          setTimeout(() => {
            location.reload()
          }, 1200)
        }
      } catch (e) {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;600&family=DM+Sans:wght@400;500;600&display=swap');

/* ── 页面 ── */
.login-page {
  position: fixed;
  inset: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #0f0f13;
  font-family: 'DM Sans', -apple-system, sans-serif;
  overflow: hidden;
}

/* ── 背景 ── */
.bg-grid {
  position: absolute;
  inset: 0;
  background-image:
    radial-gradient(circle, rgba(255,255,255,0.03) 1px, transparent 1px);
  background-size: 32px 32px;
}

.bg-glow {
  position: absolute;
  border-radius: 50%;
  filter: blur(100px);
  opacity: 0.25;
  animation: glow-drift 12s ease-in-out infinite alternate;
}

.bg-glow--1 {
  width: 500px;
  height: 500px;
  background: #dc6b2f;
  top: -15%;
  right: -10%;
}

.bg-glow--2 {
  width: 400px;
  height: 400px;
  background: #b84a1c;
  bottom: -10%;
  left: -8%;
  animation-delay: -6s;
}

@keyframes glow-drift {
  0%   { transform: translate(0, 0) scale(1); }
  100% { transform: translate(30px, -20px) scale(1.1); }
}

/* ── 卡片 ── */
.login-card {
  position: relative;
  width: 400px;
  background: rgba(22, 22, 30, 0.85);
  backdrop-filter: blur(24px);
  -webkit-backdrop-filter: blur(24px);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 16px;
  padding: 0;
  z-index: 1;
  animation: card-enter 0.6s cubic-bezier(0.16, 1, 0.3, 1) both;
  box-shadow:
    0 0 0 1px rgba(255,255,255,0.03),
    0 24px 80px -12px rgba(0,0,0,0.6);
}

@keyframes card-enter {
  from {
    opacity: 0;
    transform: translateY(24px) scale(0.97);
  }
}

/* ── 顶部装饰条 ── */
.card-accent {
  height: 3px;
  background: linear-gradient(90deg, #dc6b2f, #e8a36a, #dc6b2f);
  border-radius: 16px 16px 0 0;
}

/* ── 头部 ── */
.card-header {
  padding: 36px 36px 0;
  text-align: center;
}

.logo {
  font-family: 'JetBrains Mono', monospace;
  font-size: 22px;
  font-weight: 600;
  letter-spacing: -0.5px;
  color: #e8e4e0;
  animation: logo-enter 0.5s 0.2s cubic-bezier(0.16, 1, 0.3, 1) both;
}

@keyframes logo-enter {
  from { opacity: 0; transform: translateY(8px); }
}

.logo-bracket { color: #6b6560; }
.logo-at      { color: #dc6b2f; }
.logo-host    { color: #e8a36a; }
.logo-text    { color: #c0bdb8; }

.logo-cursor {
  color: #dc6b2f;
  animation: blink 1s step-end infinite;
}

@keyframes blink {
  50% { opacity: 0; }
}

.subtitle {
  margin: 10px 0 0;
  font-size: 13px;
  color: #6b6560;
  letter-spacing: 0.5px;
  animation: logo-enter 0.5s 0.35s cubic-bezier(0.16, 1, 0.3, 1) both;
}

/* ── 表单 ── */
.card-body {
  padding: 32px 36px 28px;
}

.input-group {
  margin-bottom: 22px;
  animation: logo-enter 0.5s 0.45s cubic-bezier(0.16, 1, 0.3, 1) both;
}

.input-group:nth-child(2) {
  animation-delay: 0.55s;
}

.input-label {
  display: block;
  font-size: 12px;
  font-weight: 500;
  color: #8a857f;
  margin-bottom: 8px;
  text-transform: uppercase;
  letter-spacing: 1px;
}

.input-wrap {
  display: flex;
  align-items: center;
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 10px;
  padding: 0 14px;
  transition: all 0.25s ease;
}

.input-wrap.is-focus {
  border-color: rgba(220, 107, 47, 0.5);
  background: rgba(220, 107, 47, 0.04);
  box-shadow: 0 0 0 3px rgba(220, 107, 47, 0.08);
}

.input-icon {
  color: #6b6560;
  display: flex;
  align-items: center;
  flex-shrink: 0;
  transition: color 0.25s ease;
}

.input-wrap.is-focus .input-icon {
  color: #dc6b2f;
}

.input-wrap input {
  flex: 1;
  background: transparent;
  border: none;
  outline: none;
  padding: 13px 12px;
  font-family: 'DM Sans', sans-serif;
  font-size: 14px;
  color: #e8e4e0;
  caret-color: #dc6b2f;
}

.input-wrap input::placeholder {
  color: #4a4540;
}

/* ── 按钮 ── */
.btn-login {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 48px;
  margin-top: 28px;
  border: none;
  border-radius: 10px;
  background: linear-gradient(135deg, #dc6b2f 0%, #c45a22 100%);
  color: #fff;
  font-family: 'DM Sans', sans-serif;
  font-size: 15px;
  font-weight: 600;
  letter-spacing: 2px;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
  animation: logo-enter 0.5s 0.65s cubic-bezier(0.16, 1, 0.3, 1) both;
}

.btn-login::before {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, transparent, rgba(255,255,255,0.12));
  opacity: 0;
  transition: opacity 0.3s ease;
}

.btn-login:hover::before {
  opacity: 1;
}

.btn-login:hover {
  transform: translateY(-1px);
  box-shadow: 0 8px 32px -4px rgba(220, 107, 47, 0.35);
}

.btn-login:active {
  transform: translateY(0);
}

.btn-login:disabled {
  cursor: not-allowed;
  opacity: 0.8;
}

/* ── Spinner ── */
.btn-spinner {
  display: inline-flex;
  align-items: center;
  gap: 10px;
}

.spinner {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(255,255,255,0.3);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* ── 成功 ── */
.card-success {
  padding: 48px 36px;
  text-align: center;
  animation: fade-in 0.4s ease both;
}

@keyframes fade-in {
  from { opacity: 0; transform: scale(0.95); }
}

.success-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: rgba(220, 107, 47, 0.1);
  border: 1px solid rgba(220, 107, 47, 0.2);
  color: #dc6b2f;
  margin-bottom: 20px;
  animation: check-pop 0.5s 0.15s cubic-bezier(0.16, 1, 0.3, 1) both;
}

@keyframes check-pop {
  from { transform: scale(0.6); opacity: 0; }
}

.success-icon svg {
  animation: check-draw 0.5s 0.3s ease both;
}

@keyframes check-draw {
  from {
    stroke-dasharray: 40;
    stroke-dashoffset: 40;
  }
  to {
    stroke-dashoffset: 0;
  }
}

.success-text {
  font-size: 18px;
  font-weight: 600;
  color: #e8e4e0;
  margin: 0;
}

.success-sub {
  font-size: 13px;
  color: #6b6560;
  margin: 6px 0 0;
}

/* ── 底部 ── */
.card-footer {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  padding: 16px 36px;
  border-top: 1px solid rgba(255, 255, 255, 0.04);
  font-family: 'JetBrains Mono', monospace;
  font-size: 11px;
  color: #4a4540;
  letter-spacing: 0.5px;
}

.footer-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #dc6b2f;
  opacity: 0.6;
}

/* ── loading 状态 ── */
.login-card.is-loading .card-body {
  opacity: 0.7;
  pointer-events: none;
}
</style>
