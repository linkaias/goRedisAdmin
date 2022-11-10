<template>
  <div>
    <el-form ref="form_son" :rules="rules" :model="formData" label-width="80px">
      <el-form-item prop="key" label="键名">
        <el-input style="width: 82%" v-model="formData.key" placeholder="请输入键名～"></el-input>
      </el-form-item>
      <el-row>
        <el-col :span="12">
          <el-form-item prop="type" label="类型">
            <el-select v-model="formData.type" placeholder="请选择数据类型">
              <el-option label="string" value="string"></el-option>
              <el-option label="list" value="list"></el-option>
              <el-option label="set" value="set"></el-option>
              <el-option label="zset" value="zset"></el-option>
              <el-option label="hash" value="hash"></el-option>
              <!--<el-option label="stream" value="stream"></el-option>-->
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="过期时间">
            <el-input v-model="formData.expire" type="number" placeholder="0为永不过期"></el-input>
          </el-form-item>
        </el-col>
      </el-row>

      <el-form-item v-if="formData.type==='zset'" prop="" label="Score">
        <el-input style="width: 82%" v-model="formData.score" placeholder="请输入score" type="number"></el-input>
      </el-form-item>
      <el-form-item v-if="formData.type==='hash'" prop="hash_key" label="键名">
        <el-input style="width: 82%" v-model="formData.hash_key" placeholder="请输入键名"></el-input>
      </el-form-item>

      <el-form-item label="值">
        <el-input type="textarea" v-model="formData.val" rows="6" placeholder="请输入值～"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="saveData" type="primary" size="small" icon="el-icon-check">保存</el-button>
        <el-button @click="close" size="small" icon="el-icon-check">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {cloneDeep} from "lodash";

export default {
  name: "formPage",
  data() {
    return {
      rules: {
        key: [
          {required: true, message: '请填写Key', trigger: 'blur'}
        ],
        hash_key: [
          {required: true, message: '请填写键名', trigger: 'blur'}
        ],
        type: [
          {required: true, message: '请选择类型', trigger: 'change'}
        ],
      },
      dbName: 0,
      formData: {}
    }
  },
  mounted() {

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
          console.log(data)
          let res = await this.$API.dbApi.reqAddVal(data)
          this.$message.success(res.message)
          this.closeAndReload()
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    }
  }
}
</script>

<style scoped>

</style>
