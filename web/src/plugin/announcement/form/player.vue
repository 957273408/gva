
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="选手名:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入选手名" />
       </el-form-item>
        <el-form-item label="战队:" prop="team">
           <el-select v-model="formData.team" placeholder="请选择战队" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in teamOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="定妆照:" prop="img">
          <SelectImage v-model="formData.img" file-type="image"/>
       </el-form-item>
        <el-form-item label="大屏定妆照:" prop="bImg">
          <SelectImage v-model="formData.bImg" file-type="image"/>
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createPlayer,
  updatePlayer,
  findPlayer
} from '@/plugin/announcement/api/player'

defineOptions({
    name: 'PlayerForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const teamOptions = ref([])
const formData = ref({
            name: '',
            team: '',
            img: "",
            bImg: "",
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPlayer({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    teamOptions.value = await getDictFunc('team')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createPlayer(formData.value)
               break
             case 'update':
               res = await updatePlayer(formData.value)
               break
             default:
               res = await createPlayer(formData.value)
               break
           }
           btnLoading.value = false
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
