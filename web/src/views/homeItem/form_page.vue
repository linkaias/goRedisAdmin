<template>
  <div class="form-page">
    <el-form ref="form_son" :rules="rules" :model="formData" label-width="80px">
      <el-form-item prop="key" :label="$t('form.keyName')">
        <el-input
          :disabled="formData.id > 0"
          v-model="formData.key"
          :placeholder="$t('form.keyPlaceholder')"
        />
      </el-form-item>

      <div class="form-row">
        <el-form-item prop="type" :label="$t('form.type')" class="form-row-item">
          <el-select :disabled="formData.id > 0" v-model="formData.type" :placeholder="$t('form.typePlaceholder')">
            <el-option label="string" value="string" />
            <el-option label="list" value="list" />
            <el-option label="set" value="set" />
            <el-option label="zset" value="zset" />
            <el-option label="hash" value="hash" />
            <el-option label="stream" value="stream" />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('form.expireTime')" class="form-row-item form-row-small">
          <el-input v-model="formData.expire" type="number" :placeholder="$t('form.expirePlaceholder')" />
        </el-form-item>
      </div>

      <el-form-item v-if="formData.type === 'zset'" :label="$t('form.score')">
        <el-input v-model="formData.score" :placeholder="$t('form.scorePlaceholder')" type="number" />
      </el-form-item>

      <el-form-item v-if="formData.type === 'hash'" prop="hash_key" :label="$t('form.hashKey')">
        <el-input v-model="formData.hash_key" :placeholder="$t('form.hashKeyPlaceholder')" />
      </el-form-item>

      <el-form-item v-if="formData.type === 'stream'" prop="stream_field" :label="$t('form.streamField')">
        <el-input v-model="formData.stream_field" :placeholder="$t('form.streamFieldPlaceholder')" />
      </el-form-item>

      <el-form-item :label="$t('form.value')">
        <el-input type="textarea" v-model="formData.val" :rows="5" :placeholder="valuePlaceholderText" />
        <div v-if="valueHintLines.length" class="value-hint">
          <div class="value-hint-label">{{ $t('form.value') }} {{ formData.type }} {{ $t('form.formatHint') }}</div>
          <code v-for="(line, i) in valueHintLines" :key="i" class="value-hint-line">{{ line }}</code>
        </div>
      </el-form-item>

      <el-form-item class="form-actions">
        <button type="button" class="btn-submit" @click="saveData">
          <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="20 6 9 17 4 12"/>
          </svg>
          {{ $t('common.save') }}
        </button>
        <button type="button" class="btn-cancel" @click="close">{{ $t('common.cancel') }}</button>
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
      dbName: 0,
      formData: {}
    }
  },
  computed: {
    rules() {
      return {
        key: [{ required: true, message: this.$t('form.keyRequired'), trigger: 'blur' }],
        hash_key: [{ required: true, message: this.$t('form.hashKeyRequired'), trigger: 'blur' }],
        stream_field: [{ required: true, message: this.$t('form.streamFieldRequired'), trigger: 'blur' }],
        type: [{ required: true, message: this.$t('form.typeRequired'), trigger: 'change' }],
      }
    },
    valuePlaceholderText() {
      const map = {
        string: 'form.valuePlaceholderString',
        list: 'form.valuePlaceholderListShort',
        set: 'form.valuePlaceholderSetShort',
        zset: 'form.valuePlaceholderZsetShort',
        hash: 'form.valuePlaceholderHash',
        stream: 'form.valuePlaceholderStreamShort',
      }
      return this.$t(map[this.formData.type] || 'form.valuePlaceholder')
    },
    valueHintLines() {
      const map = {
        list: ['item1', 'item2', 'item3'],
        set: ['member1', 'member2', 'member3'],
        zset: ['member1', 'member2'],
        stream: ['name John', 'age 25'],
      }
      return map[this.formData.type] || []
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
          stream_field: "value",
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
  align-items: flex-start;
}

.form-row-item {
  flex: 1;
  margin-bottom: 18px;
}

.form-row-item .el-select {
  width: 100%;
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

.value-hint {
  margin-top: 8px;
  padding: 10px 12px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid var(--border-subtle);
  border-radius: 6px;
}

.value-hint-label {
  font-family: var(--font-mono);
  font-size: 11px;
  color: var(--text-muted);
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin-bottom: 6px;
}

.value-hint-line {
  display: block;
  font-family: var(--font-mono);
  font-size: 12px;
  color: var(--accent-light);
  padding: 2px 0;
  line-height: 1.6;
}
</style>
