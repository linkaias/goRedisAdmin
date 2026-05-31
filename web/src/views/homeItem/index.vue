<template>
  <div class="home-page">
    <div class="home-layout">
      <!-- ── 左侧数据库列表 ── -->
      <aside class="sidebar">
        <div class="sidebar-header">
          <span class="sidebar-title">{{ $t('home.databasesTitle') }}</span>
          <button class="btn-icon" @click="getDbList()" :class="{ 'is-spinning': leftLoading }" :title="$t('common.refresh')">
            <svg width="15" height="15" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="23 4 23 10 17 10"/>
              <path d="M20.49 15a9 9 0 1 1-2.12-9.36L23 10"/>
            </svg>
          </button>
        </div>
        <div class="sidebar-body" v-loading="leftLoading">
          <div
            v-for="(item, index) in dbInfo"
            :key="index"
            class="db-item"
            :class="{ 'is-active': item.db_num === activeDb.db_num }"
            @click="handleClickLeft(item)"
          >
            <span class="db-icon">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">
                <ellipse cx="12" cy="5" rx="9" ry="3"/>
                <path d="M21 12c0 1.66-4 3-9 3s-9-1.34-9-3"/>
                <path d="M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5"/>
              </svg>
            </span>
            <span class="db-name">{{ item.show_name }}</span>
            <span class="db-count">{{ item.keys_len }}</span>
          </div>
        </div>
      </aside>

      <!-- ── 右侧内容区 ── -->
      <section class="content">
        <!-- 顶部操作栏 -->
        <div class="content-toolbar">
          <div class="toolbar-left">
            <span class="toolbar-db" v-if="activeDb.show_name">
              <span class="toolbar-db-label">{{ $t('home.current') }}</span>
              <span class="toolbar-db-name">{{ activeDb.show_name }}</span>
              <span class="toolbar-db-count">{{ activeDb.keys_len }} {{ $t('home.keyCount') }}</span>
            </span>
            <span class="toolbar-db" v-else>
              <span class="toolbar-db-label" style="color: var(--text-muted)">{{ $t('home.selectDatabase') }}</span>
            </span>
          </div>
          <div class="toolbar-right" v-if="activeDb.show_name">
            <button class="btn-tool" @click="handleClickLeft(activeDb)" :title="$t('common.refresh')">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="23 4 23 10 17 10"/>
                <path d="M20.49 15a9 9 0 1 1-2.12-9.36L23 10"/>
              </svg>
              {{ $t('common.refresh') }}
            </button>
            <button class="btn-tool btn-accent" @click="showForm({})">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="12" y1="5" x2="12" y2="19"/>
                <line x1="5" y1="12" x2="19" y2="12"/>
              </svg>
              {{ $t('home.addKey') }}
            </button>
            <button class="btn-tool btn-ghost" @click="handleFlush('db')">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="3 6 5 6 21 6"/>
                <path d="M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2"/>
              </svg>
              {{ $t('home.clearDatabase') }}
            </button>
            <el-popconfirm :title="$t('home.flushAllConfirm')" @confirm="handleFlush('all')">
              <button class="btn-tool btn-danger" slot="reference">{{ $t('home.flushAll') }}</button>
            </el-popconfirm>
          </div>
        </div>

        <!-- 数据表格区 -->
        <div class="content-body" v-loading="rightLoading">
          <template v-if="activeDb.show_name">
            <!-- 过滤栏 -->
            <div class="filter-bar">
              <div class="filter-actions">
                <button class="btn-tool btn-sm" @click="handelDelBatch">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <polyline points="3 6 5 6 21 6"/>
                    <path d="M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2"/>
                  </svg>
                  {{ $t('home.batchDelete') }}
                </button>
                <button class="btn-tool btn-sm" @click="handelExportData">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4"/>
                    <polyline points="7 10 12 15 17 10"/>
                    <line x1="12" y1="15" x2="12" y2="3"/>
                  </svg>
                  {{ $t('home.exportData') }}
                </button>
              </div>
              <div class="filter-search">
                <el-input
                  v-model="filter"
                  :placeholder="$t('home.searchPlaceholder')"
                  prefix-icon="el-icon-search"
                  @change="changeFilter"
                  clearable
                  size="small"
                />
              </div>
            </div>

            <!-- 表格 -->
            <el-table
              ref="multipleTable"
              @selection-change="handleSelectionChange"
              :data="pageData"
              style="width: 100%"
            >
              <el-table-column type="selection" width="45" />
              <el-table-column prop="id" label="ID" align="center" width="70" />
              <el-table-column prop="key" label="Key" min-width="200" />
              <el-table-column prop="expire_at" :label="$t('home.expireAt')" width="140" />
              <el-table-column prop="type" :label="$t('home.type')" width="90">
                <template v-slot:default="{row}">
                  <span class="type-badge" :class="'type-' + row.type">{{ row.type }}</span>
                </template>
              </el-table-column>
              <el-table-column prop="len" :label="$t('home.size')" width="100" />
              <el-table-column width="180" align="center" :label="$t('home.actions')">
                <template v-slot:default="{row}">
                  <div class="row-actions">
                    <button class="btn-row" :title="$t('home.view')" @click="viewData(row)">
                      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                        <circle cx="12" cy="12" r="3"/>
                      </svg>
                    </button>
                    <button class="btn-row" :title="$t('home.setExpire')" @click="expireKey(row)">
                      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="12" cy="12" r="10"/>
                        <polyline points="12 6 12 12 16 14"/>
                      </svg>
                    </button>
                    <el-popconfirm :title="$t('home.deleteKeyConfirm', { key: row.key })" @confirm="delKey(row.key)">
                      <button class="btn-row btn-row-danger" :title="$t('home.delete')" slot="reference">
                        <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                          <polyline points="3 6 5 6 21 6"/>
                          <path d="M19 6v14a2 2 0 01-2 2H7a2 2 0 01-2-2V6m3 0V4a2 2 0 012-2h4a2 2 0 012 2v2"/>
                          <line x1="10" y1="11" x2="10" y2="17"/>
                          <line x1="14" y1="11" x2="14" y2="17"/>
                        </svg>
                      </button>
                    </el-popconfirm>
                  </div>
                </template>
              </el-table-column>
            </el-table>

            <!-- 分页 -->
            <div class="pagination-bar">
              <el-pagination
                @current-change="changePage"
                background
                :page-size="limit"
                :total="total"
                :current-page="page"
                layout="prev, pager, next, ->, total"
              />
            </div>
          </template>

          <div v-else class="empty-state">
            <svg width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round">
              <ellipse cx="12" cy="5" rx="9" ry="3"/>
              <path d="M21 12c0 1.66-4 3-9 3s-9-1.34-9-3"/>
              <path d="M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5"/>
            </svg>
            <p>{{ $t('home.selectDatabaseHint') }}</p>
          </div>
        </div>
      </section>
    </div>

    <!-- ── Dialogs ── -->
    <el-dialog
      :title="$t('home.addKeyDialogTitle', { dbName: activeDb.show_name || '' })"
      :visible.sync="activeForm"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :show-close="true"
      :destroy-on-close="true"
      append-to-body
      width="520px"
    >
      <FormPage ref="p_form" @closeForm="closeForm" />
    </el-dialog>

    <el-dialog
      :title="$t('home.viewDataTitle')"
      :visible.sync="activeData"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
      :show-close="true"
      :destroy-on-close="true"
      append-to-body
      width="600px"
    >
      <DataPage ref="p_data" @closeData="closeData" />
    </el-dialog>
  </div>
</template>

<script>
import FormPage from "./form_page"
import DataPage from "./data"

export default {
  name: "homeItem",
  components: { FormPage, DataPage },
  data() {
    return {
      activeDb: {},
      dbInfo: [],
      allPageData: [],
      pageData: [],
      limit: 10,
      page: 1,
      total: 0,
      leftLoading: false,
      rightLoading: false,
      activeForm: false,
      activeData: false,
      filter: "*",
      multipleSelection: []
    }
  },
  mounted() {
    this.getDbList()
    this.getKeysByDb()
  },
  methods: {
    viewData(row) {
      this.activeData = true
      this.$nextTick(() => {
        this.$refs.p_data.initData(this.activeDb.db_num, row)
      })
    },
    handelDelBatch() {
      let waitDelKey = []
      this.multipleSelection.forEach(item => waitDelKey.push(item.key))
      if (waitDelKey.length <= 0) {
        this.$message.warning(this.$t('home.selectDeleteWarning'))
        return
      }
      this.$confirm(this.$t('home.batchDeleteConfirm'), this.$t('common.warning'), {
        confirmButtonText: this.$t('common.confirm'),
        cancelButtonText: this.$t('common.cancel'),
        type: 'warning'
      }).then(async () => {
        waitDelKey = waitDelKey.join(",")
        let res = await this.$API.dbApi.reqDelKey(this.activeDb.db_num, waitDelKey)
        if (res.code === 0) {
          this.$message.success(res.message)
          await this.reload()
        } else {
          this.$message.error(res.message)
        }
      }).catch(() => {})
    },
    handelExportData() {
      let waitDelKey = []
      this.multipleSelection.forEach(item => waitDelKey.push(item.key))
      if (waitDelKey.length <= 0) {
        this.$message.warning(this.$t('home.selectExportWarning'))
        return
      }
      this.$confirm(this.$t('home.exportConfirm'), this.$t('common.warning'), {
        confirmButtonText: this.$t('common.confirm'),
        cancelButtonText: this.$t('common.cancel'),
        type: 'warning'
      }).then(async () => {
        waitDelKey = waitDelKey.join(",")
        try {
          await this.$API.dbApi.reqExportKey(this.activeDb.db_num, waitDelKey)
          this.$message.success(this.$t('home.exportSuccess'))
        } catch (e) {
          this.$message.error(e)
        }
      }).catch(() => {})
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    expireKey(row) {
      this.$prompt(this.$t('home.expirePromptMessage'), this.$t('home.expirePromptTitle'), {
        confirmButtonText: this.$t('common.confirm'),
        cancelButtonText: this.$t('common.cancel'),
      }).then(async ({ value }) => {
        let res = await this.$API.dbApi.reqExpireKey(this.activeDb.db_num, value, row.key)
        if (res.code === 0) {
          this.$message.success(res.message)
          await this.reload()
        } else {
          this.$message.error(res.message)
        }
      }).catch(() => {})
    },
    handleEdit(row) {
      this.showForm(row)
    },
    handleFlush(type) {
      let target = type === "all"
        ? this.$t('home.flushTargetAll')
        : this.$t('home.flushTargetDb', { name: this.activeDb.show_name })
      this.$confirm(this.$t('home.flushConfirm', { target }), this.$t('common.alert'), {
        confirmButtonText: this.$t('common.confirm'),
        cancelButtonText: this.$t('common.cancel'),
        type: 'warning'
      }).then(async () => {
        let res = await this.$API.dbApi.reqFlush(type, this.activeDb.db_num)
        if (res.code === 0) {
          this.$message.success(res.message)
          await this.reload()
        } else {
          this.$message.error(res.message)
        }
      }).catch(() => {})
    },
    showForm(info) {
      this.activeForm = true
      this.$nextTick(() => {
        this.$refs.p_form.initData(this.activeDb.db_num, info)
      })
    },
    async closeForm(type) {
      this.activeForm = false
      if (type === 2) await this.reload()
    },
    async closeData() {
      this.activeData = false
    },
    async delKey(key) {
      let res = await this.$API.dbApi.reqDelKey(this.activeDb.db_num, key)
      this.$message.success(res.message)
      await this.reload()
    },
    async reload() {
      this.handleClickLeft(this.activeDb)
      let page = this.pageData.length > 1 ? this.page : this.page - 1
      this.changePage(page <= 0 ? 1 : page)
      await this.getDbList()
    },
    changePage(p) {
      this.page = p
      this.pageData = this.allPageData[p - 1 < 0 ? 0 : p - 1]
    },
    handleClickLeft(item) {
      this.activeDb = item
      this.allPageData = []
      this.pageData = []
      this.page = 1
      this.total = 0
      this.getKeysByDb(item.db_num)
    },
    changeFilter() {
      this.getKeysByDb(this.activeDb.db_num)
    },
    async getDbList() {
      this.leftLoading = true
      let res = await this.$API.dbApi.reqGetDbList()
      this.dbInfo = res.data
      if (this.dbInfo.length > 0) {
        this.activeDb = this.dbInfo[0]
      }
      this.leftLoading = false
    },
    async getKeysByDb(num) {
      this.rightLoading = true
      this.leftLoading = true
      let res = await this.$API.dbApi.reqGetKeys(num, this.filter)
      this.allPageData = []
      let info = []
      let allLength = res.data.length
      res.data.forEach((item, index) => {
        info.push(item)
        if (info.length === this.limit || index === allLength - 1) {
          this.allPageData.push(info)
          info = []
        }
      })
      this.total = allLength
      this.pageData = this.allPageData[0]
      this.rightLoading = false
      this.leftLoading = false
    }
  }
}
</script>

<style scoped>
.home-page {
  max-width: 1400px;
  margin: 0 auto;
  animation: page-enter 0.4s cubic-bezier(0.16, 1, 0.3, 1) both;
}

@keyframes page-enter {
  from { opacity: 0; transform: translateY(12px); }
}

.home-layout {
  display: flex;
  gap: 20px;
}

/* ── Sidebar ── */
.sidebar {
  width: 220px;
  flex-shrink: 0;
}

.sidebar-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 4px 12px;
}

.sidebar-title {
  font-family: var(--font-mono);
  font-size: 12px;
  font-weight: 500;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 1.5px;
}

.btn-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 30px;
  height: 30px;
  border: none;
  border-radius: 6px;
  background: transparent;
  color: var(--text-muted);
  cursor: pointer;
  transition: all 0.2s;
}

.btn-icon:hover {
  color: var(--accent);
  background: var(--accent-glow);
}

.btn-icon.is-spinning svg {
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.sidebar-body {
  background: var(--bg-card);
  border: 1px solid var(--border-subtle);
  border-radius: var(--radius);
  overflow: hidden;
  min-height: 200px;
}

.db-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 14px;
  cursor: pointer;
  transition: all 0.15s ease;
  border-left: 2px solid transparent;
}

.db-item:hover {
  background: rgba(255,255,255,0.03);
}

.db-item.is-active {
  background: var(--accent-glow);
  border-left-color: var(--accent);
}

.db-item.is-active .db-name {
  color: var(--accent-light);
}

.db-item.is-active .db-icon {
  color: var(--accent);
}

.db-icon {
  color: var(--text-muted);
  display: flex;
  flex-shrink: 0;
}

.db-name {
  flex: 1;
  font-size: 13px;
  font-weight: 500;
  color: var(--text-secondary);
}

.db-count {
  font-family: var(--font-mono);
  font-size: 11px;
  color: var(--text-muted);
  background: rgba(255,255,255,0.04);
  padding: 2px 7px;
  border-radius: 4px;
}

/* ── Content ── */
.content {
  flex: 1;
  min-width: 0;
}

.content-toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: var(--bg-card);
  border: 1px solid var(--border-subtle);
  border-radius: var(--radius);
  padding: 12px 16px;
  margin-bottom: 16px;
}

.toolbar-left {
  display: flex;
  align-items: center;
}

.toolbar-db {
  display: flex;
  align-items: center;
  gap: 8px;
}

.toolbar-db-label {
  font-size: 12px;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.toolbar-db-name {
  font-family: var(--font-mono);
  font-size: 15px;
  font-weight: 600;
  color: var(--accent-light);
}

.toolbar-db-count {
  font-family: var(--font-mono);
  font-size: 12px;
  color: var(--text-muted);
  background: rgba(255,255,255,0.04);
  padding: 2px 8px;
  border-radius: 4px;
}

.toolbar-right {
  display: flex;
  align-items: center;
  gap: 8px;
}

/* ── Buttons ── */
.btn-tool {
  display: inline-flex;
  align-items: center;
  gap: 5px;
  padding: 7px 12px;
  border: 1px solid var(--border-light);
  border-radius: 7px;
  background: transparent;
  color: var(--text-secondary);
  font-family: var(--font-body);
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}

.btn-tool:hover {
  border-color: var(--accent);
  color: var(--accent);
  background: var(--accent-glow);
}

.btn-tool.btn-accent {
  background: linear-gradient(135deg, var(--accent), #c45a22);
  border-color: transparent;
  color: #fff;
}

.btn-tool.btn-accent:hover {
  opacity: 0.9;
  box-shadow: 0 4px 16px rgba(220, 107, 47, 0.25);
}

.btn-tool.btn-ghost {
  border-color: transparent;
  color: var(--text-muted);
}

.btn-tool.btn-ghost:hover {
  color: var(--text-primary);
  background: rgba(255,255,255,0.04);
  border-color: transparent;
}

.btn-tool.btn-danger {
  border-color: var(--danger);
  color: var(--danger);
}

.btn-tool.btn-danger:hover {
  background: rgba(224, 82, 82, 0.1);
}

.btn-tool.btn-sm {
  padding: 5px 10px;
  font-size: 11px;
  border-radius: 6px;
}

/* ── Body ── */
.content-body {
  background: var(--bg-card);
  border: 1px solid var(--border-subtle);
  border-radius: var(--radius);
  padding: 16px;
  min-height: 500px;
}

/* ── Filter bar ── */
.filter-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16px;
}

.filter-actions {
  display: flex;
  gap: 6px;
}

.filter-search {
  width: 240px;
}

/* ── Type badge ── */
.type-badge {
  display: inline-block;
  padding: 2px 8px;
  border-radius: 4px;
  font-family: var(--font-mono);
  font-size: 11px;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.3px;
}

.type-string { background: rgba(76, 175, 125, 0.12); color: #4caf7d; }
.type-hash   { background: rgba(220, 107, 47, 0.12); color: #dc6b2f; }
.type-list   { background: rgba(100, 149, 237, 0.12); color: #6495ed; }
.type-set    { background: rgba(187, 134, 252, 0.12); color: #bb86fc; }
.type-zset   { background: rgba(255, 215, 0, 0.12); color: #daa520; }

/* ── Row actions ── */
.row-actions {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
}

.btn-row {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 30px;
  height: 30px;
  border: none;
  border-radius: 6px;
  background: transparent;
  color: var(--text-muted);
  cursor: pointer;
  transition: all 0.15s;
}

.btn-row:hover {
  background: rgba(255,255,255,0.06);
  color: var(--text-primary);
}

.btn-row-danger:hover {
  background: rgba(224, 82, 82, 0.1);
  color: var(--danger);
}

/* ── Pagination ── */
.pagination-bar {
  margin-top: 16px;
  display: flex;
  justify-content: center;
}

/* ── Empty state ── */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 80px 0;
  color: var(--text-muted);
}

.empty-state svg {
  margin-bottom: 16px;
  opacity: 0.4;
}

.empty-state p {
  font-size: 14px;
  margin: 0;
}
</style>
