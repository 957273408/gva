<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单类型:" prop="ddtype">
           <el-select v-model="formData.ddtype" placeholder="请选择订单类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in khtypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="姓名:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入姓名" />
       </el-form-item>
        <el-form-item label="性别:" prop="sex">
           <el-select v-model="formData.sex" placeholder="请选择性别" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="出生日期:" prop="birthday">
          <el-input v-model="formData.birthday" :clearable="true"  placeholder="请输入出生日期" />
       </el-form-item>
        <el-form-item label="收件人姓名:" prop="getuser">
          <el-input v-model="formData.getuser" :clearable="true"  placeholder="请输入收件人姓名" />
       </el-form-item>
        <el-form-item label="关系:" prop="gx">
          <el-input v-model="formData.gx" :clearable="true"  placeholder="请输入关系" />
       </el-form-item>
        <el-form-item label="安置类型:" prop="boxtype">
           <el-select v-model="formData.boxtype" placeholder="请选择安置类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in boxtypeOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="安置数量:" prop="number">
          <el-input v-model="formData.number" :clearable="true"  placeholder="请输入安置数量" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createKhinfo,
  updateKhinfo,
  findKhinfo
} from '@/api/first/khinfo'

defineOptions({
    name: 'KhinfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const khtypeOptions = ref([])
const boxtypeOptions = ref([])
const genderOptions = ref([])
const formData = ref({
            ddtype: '',
            name: '',
            sex: '',
            address: '',
            birthday: '',
            getuser: '',
            gx: '',
            boxtype: '',
            number: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findKhinfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    khtypeOptions.value = await getDictFunc('khtype')
    boxtypeOptions.value = await getDictFunc('boxtype')
    genderOptions.value = await getDictFunc('gender')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createKhinfo(formData.value)
               break
             case 'update':
               res = await updateKhinfo(formData.value)
               break
             default:
               res = await createKhinfo(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
