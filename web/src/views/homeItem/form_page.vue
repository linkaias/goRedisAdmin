<template>
  <div class="form-page">
    <el-form ref="form_son" :rules="rules" :model="formData" label-width="80px">
      <el-form-item prop="key" label="键名">
        <el-input
          :disabled="formData.id > 0"
          v-model="formData.key"
          placeholder="输入 Key 名称"
        />
      </el-form-item>

      <div class="form-row">
        <el-form-item prop="type" label="类型" class="form-row-item">
          <el-select :disabled="formData.id > 0" v-model="formData.type" placeholder="选择数据类型">
            <el-option label="string" value="string" />
            <el-option label="list" value="list" />
            <el-option label="set" value="set" />
            <el-option label="zset" value="zset" />
            <el-option label="hash" value="hash" />
          </el-select>
        </el-form-item>
        <el-form-item label="过期时间" class="form-row-item form-row-small">
          <el-input v-model="formData.expire" type="number" placeholder="0 = 永不过期" />
        </el-form-item>
      </div>

      <el-form-item v-if="formData.type === 'zset'" label="Score">
        <el-input v-model="formData.score" placeholder="输入 Score" type="number" />
      </el-form-item>

      <el-form-item v-if="formData.type === 'hash'" prop="hash_key" label="Hash Key">
        <el-input v-model="formData.hash_key" placeholder="输入 Hash 键名" />
      </el-form-item>

      <el-form-item label="值">
        <el-input type="textarea" v-model="formData.val" :rows="5" placeholder="输入值" />
      </el-form-item>

      <el-form-item class="form-actions">
        <button type="button" class="btn-submit" @click="saveData">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="20 6 9 17 4 12"/>
          </svg>
          保存
        </button>
        <button type="button" class="btn-cancel" @click="close">取消</button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {cloneDeep} from "lodash"

export default {
  name: "formPage",
  data() {
    return {
      rules: {
        key: [{ required: true, message: '请填写 Key', trigger: 'blur' }],
        hash_key: [{ required: true, message: '请填写 Hash 键名', trigger: 'blur' }],
        type: [{ required: true, message: '请选择类型', trigger: 'change' }],
      },
      dbName: 0,
      formData: {}
    }
  },
  methods: {
    initData(dbNum, info) {
      this.dbName = dbNum
      this.initFormData(info)
    },
    initFormData(info) {
      if (info.type) {
        this.formData = info
      } else {
        this.formData = {
          type: "string",
          db_num: 0,
          key: "",
          val: "",
          score: 1,
          hash_key: "",
          expire: 0,
        }
      }
    },
    close() {
      this.$emit("closeForm", 1)
    },
    closeAndReload() {
      this.$emit("closeForm", 2)
    },
    saveData() {
      this.$refs.form_son.validate(async (valid) => {
        if (valid) {
          let data = cloneDeep(this.formData)
          data.db_num = this.dbName
          data.expire = parseInt(data.expire)
          data.score = parseInt(data.score)
          let res = await this.$API.dbApi.reqAddVal(data)
          this.$message.success(res.message)
          this.closeAndReload()
        }
      })
    }
  }
}
</script>

<style scoped>
.form-page {
  animation: fade-in 0.3s ease both;
}

@keyframes fade-in {
  from { opacity: 0; transform: translateY(8px); }
}

.form-row {
  display: flex;
  gap: 12px;
}

.form-row-item {
  flex: 1;
}

.form-row-small {
  flex: 0 0 160px;
}

.form-actions {
  display: flex;
  gap: 10px;
  margin-top: 8px;
  margin-bottom: 0;
}

.btn-submit {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 9px 24px;
  border: none;
  border-radius: 8px;
  background: linear-gradient(135deg, var(--accent), #c45a22);
  color: #fff;
  font-family: var(--font-body);
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-submit:hover {
  opacity: 0.9;
  box-shadow: 0 4px 16px rgba(220, 107, 47, 0.25);
}

.btn-cancel {
  display: inline-flex;
  align-items: center;
  padding: 9px 20px;
  border: 1px solid var(--border-light);
  border-radius: 8px;
  background: transparent;
  color: var(--text-secondary);
  font-family: var(--font-body);
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-cancel:hover {
  border-color: var(--text-muted);
  color: var(--text-primary);
}
</style>
