<template>
  <div style="margin-top: 20px">
    <el-row :gutter="20">
      <el-col :span="4" :offset="2">
        <el-card>
          <el-row>
            <span style="line-height: 29px">LocalHost</span>
            <el-button @click="getDbList()" style="float: right" type="" plain size="mini"
                       icon="el-icon-refresh"></el-button>
          </el-row>
          <el-divider></el-divider>
          <ul v-loading="leftLoading" style="padding-left: 20px">
            <li @click="handleClickLeft(item)" :class="{'active_db':item.db_num===activeDb.db_num}" class="db_son"
                v-for="(item,index) in dbInfo" :key="index">
              <svg style="width: 15px;height: 15px" t="1666163640157" class="icon" viewBox="0 0 1024 1024" version="1.1"
                   xmlns="http://www.w3.org/2000/svg" p-id="1823" width="128" height="128">
                <path
                    d="M1023.786667 611.84c-0.426667 9.770667-13.354667 20.693333-39.893334 34.56-54.613333 28.458667-337.749333 144.896-397.994666 176.298667-60.288 31.402667-93.738667 31.104-141.354667 8.32-47.616-22.741333-348.842667-144.469333-403.114667-170.368-27.093333-12.970667-40.917333-23.893333-41.386666-34.218667v103.509333c0 10.325333 14.250667 21.290667 41.386666 34.261334 54.272 25.941333 355.541333 147.626667 403.114667 170.368 47.616 22.784 81.066667 23.082667 141.354667-8.362667 60.245333-31.402667 343.338667-147.797333 397.994666-176.298667 27.776-14.464 40.106667-25.728 40.106667-35.925333v-102.058667l-0.213333-0.085333z"
                    fill="" p-id="1824"></path>
                <path
                    d="M1023.744 443.093333c-0.426667 9.770667-13.354667 20.650667-39.850667 34.517334-54.613333 28.458667-337.749333 144.896-397.994666 176.298666-60.288 31.402667-93.738667 31.104-141.354667 8.362667-47.616-22.741333-348.842667-144.469333-403.114667-170.410667-27.093333-12.928-40.917333-23.893333-41.386666-34.176v103.509334c0 10.325333 14.250667 21.248 41.386666 34.218666 54.272 25.941333 355.498667 147.626667 403.114667 170.368 47.616 22.784 81.066667 23.082667 141.354667-8.32 60.245333-31.402667 343.338667-147.84 397.994666-176.298666 27.776-14.506667 40.106667-25.770667 40.106667-35.968v-102.058667l-0.256-0.042667z"
                    fill="" p-id="1825"></path>
                <path
                    d="M1023.744 268.074667c0.512-10.410667-13.098667-19.541333-40.490667-29.610667-53.248-19.498667-334.634667-131.498667-388.522666-151.253333-53.888-19.712-75.818667-18.901333-139.093334 3.84C392.234667 113.706667 92.629333 231.253333 39.338667 252.074667c-26.666667 10.496-39.68 20.181333-39.253334 30.506666V386.133333c0 10.325333 14.250667 21.248 41.386667 34.218667 54.272 25.941333 355.498667 147.669333 403.114667 170.410667 47.616 22.741333 81.066667 23.04 141.354666-8.362667 60.245333-31.402667 343.338667-147.84 397.994667-176.298667 27.776-14.506667 40.106667-25.770667 40.106667-35.968V268.074667h-0.341334zM366.72 366.08l237.269333-36.437333-71.68 105.088-165.546666-68.650667z m524.8-94.634667l-140.330667 55.466667-15.232 5.973333-140.245333-55.466666 155.392-61.44 140.373333 55.466666z m-411.989333-101.674666l-22.954667-42.325334 71.594667 27.989334 67.498666-22.101334-18.261333 43.733334 68.778667 25.770666-88.704 9.216-19.882667 47.786667-32.085333-53.290667-102.4-9.216 76.416-27.562666z m-176.768 59.733333c70.058667 0 126.805333 21.973333 126.805333 49.109333s-56.746667 49.152-126.805333 49.152-126.848-22.058667-126.848-49.152c0-27.136 56.789333-49.152 126.848-49.152z"
                    fill="" p-id="1826"></path>
              </svg>
              {{ item.show_name }} <span style="color: grey">({{ item.keys_len }})</span>
            </li>
          </ul>
        </el-card>
      </el-col>
      <el-col :span="16">
        <el-card>
        <span>
          ??????????????????<span v-show="activeDb.show_name">{{ activeDb.show_name }}  <span
            style="color: grey"> ({{ activeDb.keys_len }})</span></span>
          <span v-show="!activeDb.show_name">- - -</span>
        </span>
          <div style="float: right;">
            <el-button-group>
              <el-button @click="handleClickLeft(activeDb)" :disabled="!activeDb.show_name" size="small"
                         icon="el-icon-refresh">??????
              </el-button>
              <el-button @click="showForm({})" :disabled="!activeDb.show_name" size="small" icon="el-icon-plus">????????????
              </el-button>
              <el-button @click="handleFlush('db')" :disabled="!activeDb.show_name" size="small"
                         icon="el-icon-delete">?????????
              </el-button>
            </el-button-group>
            <el-popconfirm style="margin-left: 10px"
                           :title="'?????????????????????????????????'"
                           @confirm="handleFlush('all')"
            >
              <el-button slot="reference" type="danger" :disabled="!activeDb.show_name" size="small"
                         icon="el-icon-delete">FlushAll
              </el-button>
            </el-popconfirm>
          </div>

        </el-card>
        <el-card v-loading="rightLoading" style="margin-top: 10px;min-height: 700px">
          <el-skeleton v-show="rightLoading" :rows="12"/>
          <div v-show="activeDb.show_name">
            <el-table
                :data="pageData"
                style="width: 100%">
              <el-table-column
                  prop="id"
                  label="ID"
                  align="center"
                  width="100">
              </el-table-column>
              <el-table-column
                  prop="key"
                  label="Key"
                  width="width">
              </el-table-column>
              <el-table-column
                  prop="expire_at"
                  label="????????????"
                  width="width">
              </el-table-column>
              <el-table-column
                  prop="type"
                  label="??????"
                  width="100">
              </el-table-column>
              <el-table-column
                  width="200"
                  align="center"
                  label="??????">
                <template v-slot:default="{row}">
                  <!--                  <el-button @click="handleEdit(row)" type="primary" plain size="mini"
                                               icon="el-icon-edit-outline"></el-button>-->
                  <el-button @click="viewData(row)" type="primary" plain size="mini"
                             icon="el-icon-info"></el-button>
                  <el-popconfirm
                      :title="'???????????? ['+row.key+'] ??????'"
                      @confirm="delKey(row.key)"
                  >
                    <el-button style="margin-left: 8px" slot="reference" type="danger" plain size="mini"
                               icon="el-icon-delete"></el-button>
                  </el-popconfirm>
                </template>
              </el-table-column>
            </el-table>
            <el-pagination style="margin-top: 10px;text-align: center"
                           @current-change="changePage"
                           background
                           :page-size="limit"
                           :total="total"
                           :current-page="page"
                           layout="prev, pager, next, ->, total"
            >
            </el-pagination>
          </div>
          <el-empty v-show="!activeDb.show_name" description="????????????????????????"></el-empty>
        </el-card>
      </el-col>
    </el-row>
    <el-dialog :title="'???????????????????????????'+activeDb.show_name" :visible.sync="activeForm" :close-on-click-modal="false"
               :close-on-press-escape="false" :show-close="false" :destroy-on-close="true">
      <FormPage ref="p_form" @closeForm="closeForm"></FormPage>
    </el-dialog>

    <el-dialog title="????????????" :visible.sync="activeData" :close-on-click-modal="false"
               :close-on-press-escape="false" :show-close="false" :destroy-on-close="true">
      <DataPage ref="p_data" @closeData="closeData"></DataPage>
    </el-dialog>
  </div>
</template>
<script>
import FormPage from "./form_page"
import DataPage from "./data"

export default {
  name: "homeItem",
  components: {
    FormPage,
    DataPage
  },
  data() {
    return {
      activeDb: {},
      dbInfo: [],
      allPageData: [],
      pageData: [],
      limit: 10, //page count
      page: 1, //now page,
      total: 0,//all total
      leftLoading: false,
      rightLoading: false,
      activeForm: false,
      activeData: false,

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

    handleEdit(row) {
      this.showForm(row)
    },
    handleFlush(type) {
      let msg = type === "all" ? "???????????????" : "?????????" + this.activeDb.db_num + "?????????"
      this.$confirm('??????????????????' + msg + '??????, ?????????????', '??????', {
        confirmButtonText: '??????',
        cancelButtonText: '??????',
        type: 'warning'
      }).then(async () => {
        let res = await this.$API.dbApi.reqFlush(type, this.activeDb.db_num)
        if (res.code === 0) {
          this.$message.success(res.message)
          await this.reload()
        } else {
          this.$message.error(res.message)
        }
      }).catch(() => {
      });
    },
    showForm(info) {
      this.activeForm = true
      this.$nextTick(() => {
        this.$refs.p_form.initData(this.activeDb.db_num, info)
      })
    },
    async closeForm(type) {
      this.activeForm = false
      if (type === 2) { //????????????
        await this.reload()
      }
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
    //????????????
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
      let res = await this.$API.dbApi.reqGetKeys(num)
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
  },
}
</script>
<style scoped>
.db_son {
  list-style: none;
  padding: 10px 10px;
  cursor: pointer;
}

.active_db {
  background: #EBEEF5;
}
</style>
