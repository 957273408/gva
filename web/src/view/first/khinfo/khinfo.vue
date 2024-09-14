<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          ></el-date-picker>
        </el-form-item>

        <el-form-item label="订单类型" prop="ddtype">
          <el-select
            v-model="searchInfo.ddtype"
            clearable
            placeholder="请选择"
            @clear="
              () => {
                searchInfo.ddtype = undefined;
              }
            "
          >
            <el-option
              v-for="(item, key) in khtypeOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="姓名" prop="name">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button
            link
            type="primary"
            icon="arrow-down"
            @click="showAllQuery = true"
            v-if="!showAllQuery"
            >展开</el-button
          >
          <el-button
            link
            type="primary"
            icon="arrow-up"
            @click="showAllQuery = false"
            v-else
            >收起</el-button
          >
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-button
          icon="delete"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onDelete"
          >删除</el-button
        >
        <el-button
          icon="printer"
          style="margin-left: 10px"
          :disabled="!multipleSelection.length"
          @click="onPrint"
          >打印</el-button
        >
        <!-- <ExportTemplate  template-id="first_Khinfo" /> -->
        <!-- <ExportExcel  template-id="first_Khinfo" /> -->
        <!-- <ImportExcel  template-id="first_Khinfo" @on-success="getTableData" /> -->
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

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          label="订单类型"
          prop="ddtype"
          width="120"
        >
          <template #default="scope">
            {{ filterDict(scope.row.ddtype, khtypeOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="姓名" prop="name" width="120" />
        <el-table-column align="left" label="性别" prop="sex" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.sex, genderOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="地址" prop="address" width="120" />
        <el-table-column
          align="left"
          label="出生日期"
          prop="birthday"
          width="120"
        />
        <el-table-column
          align="left"
          label="收件人姓名"
          prop="getuser"
          width="120"
        />
        <el-table-column align="left" label="关系" prop="gx" width="120" />
        <el-table-column
          align="left"
          label="安置类型"
          prop="boxtype"
          width="120"
        >
          <template #default="scope">
            {{ filterDict(scope.row.boxtype, boxtypeOptions) }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="安置数量"
          prop="number"
          width="120"
        />
        <el-table-column
          align="left"
          label="操作"
          fixed="right"
          min-width="240"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
              ><el-icon style="margin-right: 5px"><InfoFilled /></el-icon
              >查看详情</el-button
            >
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateKhinfoFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
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
    <el-drawer
      destroy-on-close
      size="800"
      v-model="dialogFormVisible"
      :show-close="false"
      :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === "create" ? "添加" : "修改" }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form
        :model="formData"
        label-position="top"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="订单类型:" prop="ddtype">
          <el-select
            v-model="formData.ddtype"
            placeholder="请选择订单类型"
            style="width: 100%"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in khtypeOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="姓名:" prop="name">
          <el-input
            v-model="formData.name"
            :clearable="true"
            placeholder="请输入姓名"
          />
        </el-form-item>
        <el-form-item label="性别:" prop="sex">
          <el-select
            v-model="formData.sex"
            placeholder="请选择性别"
            style="width: 100%"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in genderOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input
            v-model="formData.address"
            :clearable="true"
            placeholder="请输入地址"
          />
        </el-form-item>
        <el-form-item label="出生日期:" prop="birthday">
          <el-input
            v-model="formData.birthday"
            :clearable="true"
            placeholder="请输入出生日期"
          />
        </el-form-item>
        <el-form-item label="收件人姓名:" prop="getuser">
          <el-input
            v-model="formData.getuser"
            :clearable="true"
            placeholder="请输入收件人姓名"
          />
        </el-form-item>
        <el-form-item label="关系:" prop="gx">
          <el-input
            v-model="formData.gx"
            :clearable="true"
            placeholder="请输入关系"
          />
        </el-form-item>
        <el-form-item label="安置类型:" prop="boxtype">
          <el-select
            v-model="formData.boxtype"
            placeholder="请选择安置类型"
            style="width: 100%"
            :clearable="true"
          >
            <el-option
              v-for="(item, key) in boxtypeOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="安置数量:" prop="number">
          <el-input
            v-model="formData.number"
            :clearable="true"
            placeholder="请输入安置数量"
          />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer
      destroy-on-close
      size="800"
      v-model="detailShow"
      :show-close="true"
      :before-close="closeDetailShow"
    >
      <el-descriptions :column="1" border>
        <el-descriptions-item label="订单类型">
          {{ filterDict(detailFrom.ddtype, khtypeOptions) }}
        </el-descriptions-item>
        <el-descriptions-item label="姓名">
          {{ detailFrom.name }}
        </el-descriptions-item>
        <el-descriptions-item label="性别">
          {{ filterDict(detailFrom.sex, genderOptions) }}
        </el-descriptions-item>
        <el-descriptions-item label="地址">
          {{ detailFrom.address }}
        </el-descriptions-item>
        <el-descriptions-item label="出生日期">
          {{ detailFrom.birthday }}
        </el-descriptions-item>
        <el-descriptions-item label="收件人姓名">
          {{ detailFrom.getuser }}
        </el-descriptions-item>
        <el-descriptions-item label="关系">
          {{ detailFrom.gx }}
        </el-descriptions-item>
        <el-descriptions-item label="安置类型">
          {{ filterDict(detailFrom.boxtype, boxtypeOptions) }}
        </el-descriptions-item>
        <el-descriptions-item label="安置数量">
          {{ detailFrom.number }}
        </el-descriptions-item>
      </el-descriptions>

      <el-card style="margin-top: 20px">
        <template #header>预览</template>
        <div class="pic">
          <div class="data">本命 {{ detailFrom.birthday }}</div>
          <div class="conter" v-if="detailFrom.ddtype === '1'">
            善信 {{ detailFrom.name }} 虔偹{{
              detailFrom.boxtype === "1" ? "冥箱" : "冥库"
            }}
            {{ detailFrom.number }} 只呈 {{ detailFrom.gx }}之灵
            冥中收执受用上清正一宗坛具封
          </div>
          <div class="conter" v-if="detailFrom.ddtype === '2'">
            善信 {{ detailFrom.name }} 虔偹冥箱 {{ detailFrom.number }} 只呈
            自身名下{{ detailFrom.number }}位堕胎血灵子
            冥中收执受用上清正一宗坛具封
          </div>
          <div class="address">
            {{ detailFrom.address }}
          </div>
        </div>
      </el-card>
    </el-drawer>

    <el-dialog v-model="dialogVisible" style="width: 518px" title="预览">
      <div class="printbox" id="print">
        <div
          v-for="(detailFrom, index) in multipleSelection"
          :key="index"
          class="pic"
        >
          <div class="data">本命 {{ detailFrom.birthday }}</div>
          <div class="conter" v-if="detailFrom.ddtype === '1'">
            善信 {{ detailFrom.name }} 虔偹{{
              detailFrom.boxtype === "1" ? "冥箱" : "冥库"
            }}
            {{ detailFrom.number }} 只呈 {{ detailFrom.gx }}之灵
            冥中收执受用上清正一宗坛具封
          </div>
          <div class="conter" v-if="detailFrom.ddtype === '2'">
            善信 {{ detailFrom.name }} 虔偹冥箱 {{ detailFrom.number }} 只呈
            自身名下{{ detailFrom.number }}位堕胎血灵子
            冥中收执受用上清正一宗坛具封
          </div>
          <div class="address">
            {{ detailFrom.address }}
          </div>
        </div>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="innerHtmlPrint"> 打印 </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createKhinfo,
  deleteKhinfo,
  deleteKhinfoByIds,
  updateKhinfo,
  findKhinfo,
  getKhinfoList,
} from "@/api/first/khinfo";

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
  filterDataSource,
  returnArrImg,
  onDownloadFile,
} from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";
import html2canvas from "html2canvas";
import { jsPDF } from "jspdf";

const dialogVisible = ref(false);

defineOptions({
  name: "Khinfo",
});

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false);

// 自动化生成的字典（可能为空）以及字段
const khtypeOptions = ref([]);
const boxtypeOptions = ref([]);
const genderOptions = ref([]);
const formData = ref({
  ddtype: "",
  name: "",
  sex: "",
  address: "",
  birthday: "",
  getuser: "",
  gx: "",
  boxtype: "",
  number: "",
});

// 验证规则
const rule = reactive({});

const searchRule = reactive({
  createdAt: [
    {
      validator: (_rule, _value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
});

const elFormRef = ref();
const elSearchFormRef = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    page.value = 1;
    pageSize.value = 10;
    getTableData();
  });
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getKhinfoList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
  khtypeOptions.value = await getDictFunc("khtype");
  boxtypeOptions.value = await getDictFunc("boxtype");
  genderOptions.value = await getDictFunc("gender");
};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteKhinfoFunc(row);
  });
};

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(async () => {
    const IDs = [];
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: "warning",
        message: "请选择要删除的数据",
      });
      return;
    }
    multipleSelection.value &&
      multipleSelection.value.map((item) => {
        IDs.push(item.ID);
      });
    const res = await deleteKhinfoByIds({ IDs });
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "删除成功",
      });
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--;
      }
      getTableData();
    }
  });
};

const onPrint = () => {
  if (multipleSelection.value.length !== 4) {
    ElMessage({
      type: "info",
      message: "请勾选3项进行打印",
    });
    return;
  }
  window.open('http://124.221.34.24:9998?ids='+ multipleSelection.value.map(item=>item.ID).join(','))
};

const innerHtmlPrint = () => {
  var target = document.getElementById("print");
  target.style.background = "#FFFFFF";

  html2canvas(target, { scale: 3 }).then(function (canvas) {

    // var pdf = new JSPDF('p', 'mm', 'a4'); // A4纸，纵向
    var direction = "p"; // 默认纵向打印
    var size = "a3";
    var printW = 297;
    var printH = 420;
    if (printW && printH) {
      size = [printW, printH];
      if (printW <= printH) {
        direction = "p";
      } else {
        direction = "l";
      }
    }
    var pdf = new jsPDF(direction, "mm", size); // A4纸，纵向
    var ctx = canvas.getContext("2d");
    var a4w = printW ? printW : 297;
    var a4h = printH ? printH : 420; // A4大小，210mm x 297mm，四边各保留10mm的边距，显示区域190x277
    var imgHeight = Math.floor((a4h * canvas.width) / a4w); // 按A4显示比例换算一页图像的像素高度
    var renderedHeight = 0;

    while (renderedHeight < canvas.height) {
      var page = document.createElement("canvas");
      page.width = canvas.width;
      page.height = Math.min(imgHeight, canvas.height - renderedHeight); // 可能内容不足一页

      // 用getImageData剪裁指定区域，并画到前面建立的canvas对象中
      page
        .getContext("2d")
        .putImageData(
          ctx.getImageData(
            0,
            renderedHeight,
            canvas.width,
            Math.min(imgHeight, canvas.height - renderedHeight)
          ),
          0,
          0
        );

      pdf.addImage(
        page.toDataURL("image/jpeg", 1.0),
        "JPEG",
        10,
        10,
        a4w,
        Math.min(a4h, (a4w * page.height) / page.width)
      ); // 添加图像到页面，保留10mm边距

      renderedHeight += imgHeight;
      if (renderedHeight < canvas.height) {
        pdf.addPage();
      } // 若是后面还有内容，添加一个空页
      page.remove();
    }
    const link = window.URL.createObjectURL(pdf.output("blob"));
    window.open(link);
  });
  dialogVisible.value = false;
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateKhinfoFunc = async (row) => {
  const res = await findKhinfo({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteKhinfoFunc = async (row) => {
  const res = await deleteKhinfo({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    ddtype: "",
    name: "",
    sex: "",
    address: "",
    birthday: "",
    getuser: "",
    gx: "",
    boxtype: "",
    number: "",
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createKhinfo(formData.value);
        break;
      case "update":
        res = await updateKhinfo(formData.value);
        break;
      default:
        res = await createKhinfo(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};

const detailFrom = ref({});

// 查看详情控制标记
const detailShow = ref(false);

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true;
};

// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findKhinfo({ ID: row.ID });
  if (res.code === 0) {
    detailFrom.value = res.data;
    openDetailShow();
  }
};

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false;
  detailFrom.value = {};
};
</script>

<style>
@font-face {
  font-family: "STLITI";
  src: url("../../../assets/font/STLITI.ttf") format("truetype");
  font-weight: normal;
  font-style: normal;
}

.printbox {
  display: flex;
  .pic {
    border: none;
    width: 150px;
    height: 840px;
    padding: 20px 0 0 0;
    > div {
      font-weight: 400;
      width: 50px;
      font-size: 15px;
      color: #231815 100%;
      letter-spacing: -1px;
      display: flex;
      align-items: center;
      -webkit-font-smoothing: antialiased;
      -moz-osx-font-smoothing: grayscale;
    }
  }
}

.pic {
  border: 1px dashed #000;
  display: flex;
  width: 74.25px;
  height: 420px;
  padding: 10.5px 0 0 0;
  justify-content: center;
  > div {
    font-family: "STLITI";
    writing-mode: vertical-rl;
    font-weight: 400;
    width: 25px;
    font-size: 11px;
    color: #231815 100%;
    letter-spacing: -1px;
    display: flex;
    align-items: center;
  }
}
</style>
