
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
         <el-date-picker
                  v-model="searchInfo.createdAtRange"
                  class="w-[380px]"
                  type="datetimerange"
                  range-separator="至"
                  start-placeholder="开始时间"
                  end-placeholder="结束时间"
                />
       </el-form-item>
      
            <el-form-item label="赛程ID" prop="scheduleId">
  <el-input v-model="searchInfo.scheduleId" placeholder="搜索条件" />
</el-form-item>
           
            <el-form-item label="队伍名" prop="teamId">
  <el-select v-model="searchInfo.teamId" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.teamId=undefined}">
    <el-option v-for="(item,key) in teamOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
           
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="赛程ID" prop="scheduleId" width="120" />

            <el-table-column align="left" label="队伍名" prop="teamId" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.teamId,teamOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="小局" prop="round" width="120" />

            <el-table-column align="left" label="红蓝" prop="sort" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.sort,SortOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="Ban1" prop="ban1" width="120" />

            <el-table-column align="left" label="Ban2" prop="ban2" width="120" />

            <el-table-column align="left" label="Ban3" prop="ban3" width="120" />

            <el-table-column align="left" label="Ban4" prop="ban4" width="120" />

            <el-table-column align="left" label="ban5" prop="ban5" width="120" />

            <el-table-column align="left" label="选手1" prop="player1" width="120" />

            <el-table-column align="left" label="英雄1" prop="hero1" width="120" />

            <el-table-column align="left" label="召唤师技能1" prop="heroSkill1" width="120" />

            <el-table-column align="left" label="选手2" prop="player2" width="120" />

            <el-table-column align="left" label="英雄2" prop="hero2" width="120" />

            <el-table-column align="left" label="召唤师技能2" prop="heroSkill2" width="120" />

            <el-table-column align="left" label="选手3" prop="player3" width="120" />

            <el-table-column align="left" label="英雄3" prop="hero3" width="120" />

            <el-table-column align="left" label="召唤师技能3" prop="heroSkill3" width="120" />

            <el-table-column align="left" label="选手4" prop="player4" width="120" />

            <el-table-column align="left" label="英雄4" prop="hero4" width="120" />

            <el-table-column align="left" label="召唤师技能4" prop="heroSkill4" width="120" />

            <el-table-column align="left" label="选手5" prop="player5" width="120" />

            <el-table-column align="left" label="英雄5" prop="hero5" width="120" />

            <el-table-column align="left" label="召唤师技能5" prop="heroSkill5" width="120" />

        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateBplistFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
             <el-form-item label="赛程ID:" prop="scheduleId">
    <el-input v-model="formData.scheduleId" :clearable="true" placeholder="请输入赛程ID" />
</el-form-item>
             <el-form-item label="队伍名:" prop="teamId">
    <el-select v-model="formData.teamId" placeholder="请选择队伍名" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in teamOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="小局:" prop="round">
    <el-input v-model.number="formData.round" :clearable="true" placeholder="请输入小局" />
</el-form-item>
             <el-form-item label="红蓝:" prop="sort">
    <el-select v-model="formData.sort" placeholder="请选择红蓝" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in SortOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
             <el-form-item label="Ban1:" prop="ban1">
    <el-input v-model="formData.ban1" :clearable="true" placeholder="请输入Ban1" />
</el-form-item>
             <el-form-item label="Ban2:" prop="ban2">
    <el-input v-model="formData.ban2" :clearable="true" placeholder="请输入Ban2" />
</el-form-item>
             <el-form-item label="Ban3:" prop="ban3">
    <el-input v-model="formData.ban3" :clearable="true" placeholder="请输入Ban3" />
</el-form-item>
             <el-form-item label="Ban4:" prop="ban4">
    <el-input v-model="formData.ban4" :clearable="true" placeholder="请输入Ban4" />
</el-form-item>
             <el-form-item label="ban5:" prop="ban5">
    <el-input v-model="formData.ban5" :clearable="true" placeholder="请输入ban5" />
</el-form-item>
             <el-form-item label="选手1:" prop="player1">
    <el-input v-model="formData.player1" :clearable="true" placeholder="请输入选手1" />
</el-form-item>
             <el-form-item label="英雄1:" prop="hero1">
    <el-input v-model="formData.hero1" :clearable="true" placeholder="请输入英雄1" />
</el-form-item>
             <el-form-item label="召唤师技能1:" prop="heroSkill1">
    <el-input v-model="formData.heroSkill1" :clearable="true" placeholder="请输入召唤师技能1" />
</el-form-item>
             <el-form-item label="选手2:" prop="player2">
    <el-input v-model="formData.player2" :clearable="true" placeholder="请输入选手2" />
</el-form-item>
             <el-form-item label="英雄2:" prop="hero2">
    <el-input v-model="formData.hero2" :clearable="true" placeholder="请输入英雄2" />
</el-form-item>
             <el-form-item label="召唤师技能2:" prop="heroSkill2">
    <el-input v-model="formData.heroSkill2" :clearable="true" placeholder="请输入召唤师技能2" />
</el-form-item>
             <el-form-item label="选手3:" prop="player3">
    <el-input v-model="formData.player3" :clearable="true" placeholder="请输入选手3" />
</el-form-item>
             <el-form-item label="英雄3:" prop="hero3">
    <el-input v-model="formData.hero3" :clearable="true" placeholder="请输入英雄3" />
</el-form-item>
             <el-form-item label="召唤师技能3:" prop="heroSkill3">
    <el-input v-model="formData.heroSkill3" :clearable="true" placeholder="请输入召唤师技能3" />
</el-form-item>
             <el-form-item label="选手4:" prop="player4">
    <el-input v-model="formData.player4" :clearable="true" placeholder="请输入选手4" />
</el-form-item>
             <el-form-item label="英雄4:" prop="hero4">
    <el-input v-model="formData.hero4" :clearable="true" placeholder="请输入英雄4" />
</el-form-item>
             <el-form-item label="召唤师技能4:" prop="heroSkill4">
    <el-input v-model="formData.heroSkill4" :clearable="true" placeholder="请输入召唤师技能4" />
</el-form-item>
             <el-form-item label="选手5:" prop="player5">
    <el-input v-model="formData.player5" :clearable="true" placeholder="请输入选手5" />
</el-form-item>
             <el-form-item label="英雄5:" prop="hero5">
    <el-input v-model="formData.hero5" :clearable="true" placeholder="请输入英雄5" />
</el-form-item>
             <el-form-item label="召唤师技能5:" prop="heroSkill5">
    <el-input v-model="formData.heroSkill5" :clearable="true" placeholder="请输入召唤师技能5" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                 <el-descriptions-item label="赛程ID">
    {{ detailForm.scheduleId }}
</el-descriptions-item>
                 <el-descriptions-item label="队伍名">
    {{ detailForm.teamId }}
</el-descriptions-item>
                 <el-descriptions-item label="小局">
    {{ detailForm.round }}
</el-descriptions-item>
                 <el-descriptions-item label="红蓝">
    {{ detailForm.sort }}
</el-descriptions-item>
                 <el-descriptions-item label="Ban1">
    {{ detailForm.ban1 }}
</el-descriptions-item>
                 <el-descriptions-item label="Ban2">
    {{ detailForm.ban2 }}
</el-descriptions-item>
                 <el-descriptions-item label="Ban3">
    {{ detailForm.ban3 }}
</el-descriptions-item>
                 <el-descriptions-item label="Ban4">
    {{ detailForm.ban4 }}
</el-descriptions-item>
                 <el-descriptions-item label="ban5">
    {{ detailForm.ban5 }}
</el-descriptions-item>
                 <el-descriptions-item label="选手1">
    {{ detailForm.player1 }}
</el-descriptions-item>
                 <el-descriptions-item label="英雄1">
    {{ detailForm.hero1 }}
</el-descriptions-item>
                 <el-descriptions-item label="召唤师技能1">
    {{ detailForm.heroSkill1 }}
</el-descriptions-item>
                 <el-descriptions-item label="选手2">
    {{ detailForm.player2 }}
</el-descriptions-item>
                 <el-descriptions-item label="英雄2">
    {{ detailForm.hero2 }}
</el-descriptions-item>
                 <el-descriptions-item label="召唤师技能2">
    {{ detailForm.heroSkill2 }}
</el-descriptions-item>
                 <el-descriptions-item label="选手3">
    {{ detailForm.player3 }}
</el-descriptions-item>
                 <el-descriptions-item label="英雄3">
    {{ detailForm.hero3 }}
</el-descriptions-item>
                 <el-descriptions-item label="召唤师技能3">
    {{ detailForm.heroSkill3 }}
</el-descriptions-item>
                 <el-descriptions-item label="选手4">
    {{ detailForm.player4 }}
</el-descriptions-item>
                 <el-descriptions-item label="英雄4">
    {{ detailForm.hero4 }}
</el-descriptions-item>
                 <el-descriptions-item label="召唤师技能4">
    {{ detailForm.heroSkill4 }}
</el-descriptions-item>
                 <el-descriptions-item label="选手5">
    {{ detailForm.player5 }}
</el-descriptions-item>
                 <el-descriptions-item label="英雄5">
    {{ detailForm.hero5 }}
</el-descriptions-item>
                 <el-descriptions-item label="召唤师技能5">
    {{ detailForm.heroSkill5 }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createBplist,
  deleteBplist,
  deleteBplistByIds,
  updateBplist,
  findBplist,
  getBplistList
} from '@/plugin/announcement/api/bplist'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'Bplist'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getBplistList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    teamOptions.value = await getDictFunc('team')
    SortOptions.value = await getDictFunc('Sort')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteBplistFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteBplistByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateBplistFunc = async(row) => {
    const res = await findBplist({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteBplistFunc = async (row) => {
    const res = await deleteBplist({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findBplist({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
