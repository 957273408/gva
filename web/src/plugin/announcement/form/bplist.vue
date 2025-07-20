
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="赛程ID:" prop="scheduleId">
          <el-input v-model="formData.scheduleId" :clearable="true"  placeholder="请输入赛程ID" />
       </el-form-item>
        <el-form-item label="队伍名:" prop="teamId">
           <el-select v-model="formData.teamId" placeholder="请选择队伍名" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in teamOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="小局:" prop="round">
          <el-input v-model.number="formData.round" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="红蓝:" prop="sort">
           <el-select v-model="formData.sort" placeholder="请选择红蓝" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in SortOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="Ban1:" prop="ban1">
          <el-input v-model="formData.ban1" :clearable="true"  placeholder="请输入Ban1" />
       </el-form-item>
        <el-form-item label="Ban2:" prop="ban2">
          <el-input v-model="formData.ban2" :clearable="true"  placeholder="请输入Ban2" />
       </el-form-item>
        <el-form-item label="Ban3:" prop="ban3">
          <el-input v-model="formData.ban3" :clearable="true"  placeholder="请输入Ban3" />
       </el-form-item>
        <el-form-item label="Ban4:" prop="ban4">
          <el-input v-model="formData.ban4" :clearable="true"  placeholder="请输入Ban4" />
       </el-form-item>
        <el-form-item label="ban5:" prop="ban5">
          <el-input v-model="formData.ban5" :clearable="true"  placeholder="请输入ban5" />
       </el-form-item>
        <el-form-item label="选手1:" prop="player1">
          <el-input v-model="formData.player1" :clearable="true"  placeholder="请输入选手1" />
       </el-form-item>
        <el-form-item label="英雄1:" prop="hero1">
          <el-input v-model="formData.hero1" :clearable="true"  placeholder="请输入英雄1" />
       </el-form-item>
        <el-form-item label="召唤师技能1:" prop="heroSkill1">
          <el-input v-model="formData.heroSkill1" :clearable="true"  placeholder="请输入召唤师技能1" />
       </el-form-item>
        <el-form-item label="选手2:" prop="player2">
          <el-input v-model="formData.player2" :clearable="true"  placeholder="请输入选手2" />
       </el-form-item>
        <el-form-item label="英雄2:" prop="hero2">
          <el-input v-model="formData.hero2" :clearable="true"  placeholder="请输入英雄2" />
       </el-form-item>
        <el-form-item label="召唤师技能2:" prop="heroSkill2">
          <el-input v-model="formData.heroSkill2" :clearable="true"  placeholder="请输入召唤师技能2" />
       </el-form-item>
        <el-form-item label="选手3:" prop="player3">
          <el-input v-model="formData.player3" :clearable="true"  placeholder="请输入选手3" />
       </el-form-item>
        <el-form-item label="英雄3:" prop="hero3">
          <el-input v-model="formData.hero3" :clearable="true"  placeholder="请输入英雄3" />
       </el-form-item>
        <el-form-item label="召唤师技能3:" prop="heroSkill3">
          <el-input v-model="formData.heroSkill3" :clearable="true"  placeholder="请输入召唤师技能3" />
       </el-form-item>
        <el-form-item label="选手4:" prop="player4">
          <el-input v-model="formData.player4" :clearable="true"  placeholder="请输入选手4" />
       </el-form-item>
        <el-form-item label="英雄4:" prop="hero4">
          <el-input v-model="formData.hero4" :clearable="true"  placeholder="请输入英雄4" />
       </el-form-item>
        <el-form-item label="召唤师技能4:" prop="heroSkill4">
          <el-input v-model="formData.heroSkill4" :clearable="true"  placeholder="请输入召唤师技能4" />
       </el-form-item>
        <el-form-item label="选手5:" prop="player5">
          <el-input v-model="formData.player5" :clearable="true"  placeholder="请输入选手5" />
       </el-form-item>
        <el-form-item label="英雄5:" prop="hero5">
          <el-input v-model="formData.hero5" :clearable="true"  placeholder="请输入英雄5" />
       </el-form-item>
        <el-form-item label="召唤师技能5:" prop="heroSkill5">
          <el-input v-model="formData.heroSkill5" :clearable="true"  placeholder="请输入召唤师技能5" />
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
  createBplist,
  updateBplist,
  findBplist
} from '@/plugin/announcement/api/bplist'

defineOptions({
    name: 'BplistForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const teamOptions = ref([])
const SortOptions = ref([])
const formData = ref({
            scheduleId: '',
            teamId: '',
            round: undefined,
            sort: '',
            ban1: '',
            ban2: '',
            ban3: '',
            ban4: '',
            ban5: '',
            player1: '',
            hero1: '',
            heroSkill1: '',
            player2: '',
            hero2: '',
            heroSkill2: '',
            player3: '',
            hero3: '',
            heroSkill3: '',
            player4: '',
            hero4: '',
            heroSkill4: '',
            player5: '',
            hero5: '',
            heroSkill5: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBplist({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    teamOptions.value = await getDictFunc('team')
    SortOptions.value = await getDictFunc('Sort')
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
               res = await createBplist(formData.value)
               break
             case 'update':
               res = await updateBplist(formData.value)
               break
             default:
               res = await createBplist(formData.value)
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
