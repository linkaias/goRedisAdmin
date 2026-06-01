<template>
  <div class="app-root">
    <el-container v-if="!isLogin" class="app-container">
      <el-header class="app-header" height="60px">
        <HeaderItem />
      </el-header>
      <el-main class="app-main">
        <router-view />
      </el-main>
    </el-container>
    <router-view v-else />
  </div>
</template>

<script>
import HeaderItem from "@/views/header"
import {GetToken} from "@/utils/token"

export default {
  name: 'App',
  components: {HeaderItem},
  computed: {
    isLogin() {
      return !GetToken()
    }
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;600&family=DM+Sans:wght@400;500;600&display=swap');

:root {
  --bg-base: #0f0f13;
  --bg-surface: rgba(22, 22, 30, 0.85);
  --bg-card: rgba(28, 28, 38, 0.7);
  --bg-hover: rgba(220, 107, 47, 0.06);
  --bg-input: rgba(255, 255, 255, 0.04);
  --border-subtle: rgba(255, 255, 255, 0.06);
  --border-light: rgba(255, 255, 255, 0.1);
  --text-primary: #e8e4e0;
  --text-secondary: #8a857f;
  --text-muted: #5a5550;
  --accent: #dc6b2f;
  --accent-light: #e8a36a;
  --accent-glow: rgba(220, 107, 47, 0.15);
  --danger: #e05252;
  --success: #4caf7d;
  --radius: 10px;
  --font-body: 'DM Sans', -apple-system, sans-serif;
  --font-mono: 'JetBrains Mono', monospace;
}

* {
  box-sizing: border-box;
}

body {
  margin: 0;
  padding: 0;
  background: var(--bg-base);
  font-family: var(--font-body);
  color: var(--text-primary);
  -webkit-font-smoothing: antialiased;
}

.app-root {
  min-height: 100vh;
}

.app-container {
  min-height: 100vh;
}

/* ── Header ── */
.app-header {
  padding: 0 24px;
  background: var(--bg-surface);
  border-bottom: 1px solid var(--border-subtle);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  position: sticky;
  top: 0;
  z-index: 100;
}

/* ── Main ── */
.app-main {
  padding: 24px;
  min-height: calc(100vh - 60px);
  background:
    radial-gradient(circle at 70% 10%, rgba(220, 107, 47, 0.03) 0%, transparent 50%),
    var(--bg-base);
}

/* ── Element UI 暗色覆盖 ── */
.el-card {
  background: var(--bg-card) !important;
  border: 1px solid var(--border-subtle) !important;
  border-radius: var(--radius) !important;
  color: var(--text-primary) !important;
  box-shadow: 0 4px 24px rgba(0,0,0,0.2) !important;
}

.el-card__header {
  border-bottom: 1px solid var(--border-subtle) !important;
  padding: 16px 20px !important;
  color: var(--text-primary) !important;
}

.el-card__body {
  color: var(--text-primary) !important;
}

/* ── Table ── */
.el-table {
  background-color: transparent !important;
  color: var(--text-primary) !important;
}

.el-table th.el-table__cell {
  background-color: rgba(255,255,255,0.03) !important;
  color: var(--text-secondary) !important;
  border-bottom: 1px solid var(--border-subtle) !important;
  font-family: var(--font-mono);
  font-size: 12px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.el-table td.el-table__cell {
  border-bottom: 1px solid var(--border-subtle) !important;
  font-size: 13px;
}

.el-table tr {
  background-color: transparent !important;
}

.el-table--enable-row-hover .el-table__body tr:hover > td.el-table__cell {
  background-color: var(--bg-hover) !important;
}

.el-table__empty-block {
  background: transparent !important;
}

.el-table__empty-text {
  color: var(--text-muted) !important;
}

.el-table::before,
.el-table--border::after {
  display: none !important;
}

.el-table--border .el-table__cell {
  border-right: none !important;
}

.el-table__cell {
  border-right: none !important;
}

.el-table--scrollable-x .el-table__body-wrapper {
  overflow-x: hidden;
}

/* ── Checkbox (table) ── */
.el-checkbox__inner {
  background: transparent !important;
  border-color: var(--border-light) !important;
}

.el-checkbox__input.is-checked .el-checkbox__inner {
  background-color: var(--accent) !important;
  border-color: var(--accent) !important;
}

/* ── Pagination ── */
.el-pagination {
  font-family: var(--font-mono) !important;
}

.el-pagination button {
  background: transparent !important;
  color: var(--text-secondary) !important;
}

.el-pager li {
  background: transparent !important;
  color: var(--text-secondary) !important;
  border-radius: 6px !important;
}

.el-pager li.active {
  color: var(--accent) !important;
}

.el-pagination .el-pagination__total {
  color: var(--text-muted) !important;
}

/* ── Dialog ── */
.el-dialog__wrapper {
  z-index: 3000 !important;
  background: rgba(0, 0, 0, 0.6);
}

.el-dialog {
  background: #1c1c26 !important;
  border: 1px solid var(--border-light) !important;
  border-radius: var(--radius) !important;
  box-shadow: 0 24px 80px rgba(0,0,0,0.7) !important;
  margin-top: 8vh !important;
  position: relative;
  z-index: 3001;
}

.el-dialog__header {
  border-bottom: 1px solid var(--border-subtle) !important;
  padding: 16px 24px !important;
}

.el-dialog__title {
  color: var(--text-primary) !important;
  font-family: var(--font-mono);
  font-size: 14px !important;
}

.el-dialog__headerbtn .el-dialog__close {
  color: var(--text-muted) !important;
}

.el-dialog__body {
  color: var(--text-primary) !important;
  padding: 24px !important;
}

.el-dialog__footer {
  border-top: 1px solid var(--border-subtle) !important;
}

/* ── Form ── */
.el-form-item__label {
  color: var(--text-secondary) !important;
  font-size: 12px !important;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.el-input__inner,
.el-textarea__inner {
  background: var(--bg-input) !important;
  border: 1px solid var(--border-subtle) !important;
  color: var(--text-primary) !important;
  border-radius: 8px !important;
  font-family: var(--font-body) !important;
  caret-color: var(--accent);
}

.el-input__inner:focus,
.el-textarea__inner:focus {
  border-color: rgba(220, 107, 47, 0.5) !important;
  box-shadow: 0 0 0 3px rgba(220, 107, 47, 0.08) !important;
}

.el-input__inner::placeholder,
.el-textarea__inner::placeholder {
  color: var(--text-muted) !important;
}

.el-input-group__append,
.el-input-group__prepend {
  background: rgba(255,255,255,0.06) !important;
  border-color: var(--border-subtle) !important;
  color: var(--text-secondary) !important;
}

/* ── Select dropdown ── */
.el-select-dropdown {
  background: var(--bg-card) !important;
  border: 1px solid var(--border-subtle) !important;
  z-index: 3100 !important;
}

.el-select-dropdown__item {
  color: var(--text-primary) !important;
}

.el-select-dropdown__item.hover,
.el-select-dropdown__item:hover {
  background: var(--bg-hover) !important;
}

.el-select-dropdown__item.selected {
  color: var(--accent) !important;
}

.el-select .el-input .el-select__caret {
  color: var(--text-muted) !important;
}

.el-dialog .el-select {
  width: 100%;
}

/* ── Button ── */
.el-button--default {
  background: transparent !important;
  border-color: var(--border-light) !important;
  color: var(--text-secondary) !important;
}

.el-button--default:hover {
  border-color: var(--accent) !important;
  color: var(--accent) !important;
  background: var(--accent-glow) !important;
}

.el-button--primary {
  background: linear-gradient(135deg, var(--accent), #c45a22) !important;
  border: none !important;
  color: #fff !important;
}

.el-button--primary:hover {
  opacity: 0.9;
}

.el-button--danger {
  background: transparent !important;
  border-color: var(--danger) !important;
  color: var(--danger) !important;
}

.el-button--danger:hover {
  background: rgba(224, 82, 82, 0.1) !important;
}

.el-button--mini {
  border-radius: 6px !important;
  font-size: 12px !important;
}

/* ── Popconfirm ── */
.el-popconfirm__action {
  text-align: right;
}

.el-popover {
  background: var(--bg-card) !important;
  border: 1px solid var(--border-subtle) !important;
  color: var(--text-primary) !important;
}

/* ── Divider ── */
.el-divider {
  background-color: var(--border-subtle) !important;
}

.el-divider__text {
  background: var(--bg-card) !important;
  color: var(--text-muted) !important;
}

/* ── Message Box ── */
.el-message-box {
  background: var(--bg-card) !important;
  border: 1px solid var(--border-subtle) !important;
  border-radius: var(--radius) !important;
}

.el-message-box__title {
  color: var(--text-primary) !important;
}

.el-message-box__message {
  color: var(--text-secondary) !important;
}

.el-message-box__headerbtn .el-message-box__close {
  color: var(--text-muted) !important;
}

/* ── Message Toast ── */
.el-message {
  background: var(--bg-card) !important;
  border: 1px solid var(--border-subtle) !important;
  border-radius: 8px !important;
}

.el-message__content {
  color: var(--text-primary) !important;
}

/* ── Loading ── */
.el-loading-mask {
  background-color: rgba(15, 15, 19, 0.7) !important;
  z-index: 1000 !important;
}

.el-loading-spinner .circular {
  stroke: var(--accent) !important;
}

/* ── Empty ── */
.el-empty__description p {
  color: var(--text-muted) !important;
}

/* ── Skeleton ── */
.el-skeleton__item {
  background: rgba(255,255,255,0.04) !important;
}

/* ── Popover (global) ── */
.el-popconfirm__main {
  color: var(--text-secondary) !important;
}

/* ── Menu override ── */
.el-menu {
  background-color: transparent !important;
  border: none !important;
}

.el-menu--horizontal > .el-menu-item {
  color: var(--text-secondary) !important;
  border-bottom: none !important;
  height: 60px;
  line-height: 60px;
  font-size: 13px;
  font-weight: 500;
  transition: color 0.2s;
}

.el-menu--horizontal > .el-menu-item:hover,
.el-menu--horizontal > .el-menu-item.is-active {
  color: var(--accent-light) !important;
  background: transparent !important;
}

.el-menu--horizontal > .el-menu-item:not(.is-disabled):focus {
  background: transparent !important;
}

/* ── Scrollbar ── */
::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: rgba(255,255,255,0.08);
  border-radius: 3px;
}

::-webkit-scrollbar-thumb:hover {
  background: rgba(255,255,255,0.15);
}
</style>
