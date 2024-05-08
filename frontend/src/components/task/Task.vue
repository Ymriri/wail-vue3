<template>
  <el-form :inline="true" :model="tasksPageRequest">
    <el-form-item label="任务名称">
      <el-input v-model="tasksPageRequest.taskName" placeholder="请输入任务名称" class="form-input" clearable/>
    </el-form-item>
    <el-form-item>
      <!--      <el-select-->
      <!--          v-model="goodsPageRequest.goodsType"-->
      <!--          placeholder="请选择类型"-->
      <!--          clearable-->
      <!--          class="form-select"-->
      <!--      >-->
      <el-checkbox v-model="noAccept" label="任务未完成" size="large"/>
      <!--        <el-option v-for="item in goodsTypeOpts" :key="item.id" :label="item.name" :value="item.name"/>-->
      <!--      </el-select>-->
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">查询</el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="success" @click="addGoods">新增</el-button>
    </el-form-item>
  </el-form>
  <el-table :data="tablePageVOList" stripe style="width: 100%">
    <el-table-column prop="id" label="ID"/>
    <el-table-column prop="taskName" label="任务名称"/>
    <el-table-column prop="taskDescription" label="任务描述"/>
    <el-table-column prop="accessPath" label="ftp访问路径"/>
    <el-table-column prop="taskStartTime" label="开始时间"/>
    <el-table-column prop="taskDeadline" label="截止日期"/>
    <el-table-column prop="taskEndTime" label="任务终止时间"/>
    <el-table-column prop="taskStatus" label="任务状态"/>
    <el-table-column label="操作">
      <template #default="scope">
        <el-button type="primary" :icon="Edit" circle/>
        <el-button type="danger" :icon="Delete" circle @click="deleteGoods(scope.row.id)"/>
      </template>
    </el-table-column>
  </el-table>
  <!-- 如果需要根据tableData的实际长度来控制分页组件 -->
  <div class="example-pagination-block" v-if="tablePageVOList.length > 0">
    <el-pagination
        layout="prev, pager, next"
        :total="pageTotal"
        @current-change="handleCurrentChange"
    />
  </div>
  <!--  商品编辑对话框  -->
  <el-dialog
      v-model="dialogVisible"
      title="编辑商品"
      :before-close="handleClose"
  >
    <GoodsEdit
        :goods-v-o="currentGoodsVO"
        @refresh="onSubmit"
        @close="closeDialog"
    />
  </el-dialog>
  <!--  商品删除对话框  -->
  <el-dialog
      v-model="deleteDialogVisible"
      title="删除任务"
      :before-close="handleClose"
  >
    <GoodsDeleteDialog
        :goods-id="currentGoodsIdToDelete"
        @refresh="onSubmit"
        @close="closeDeleteDialog"
    />
  </el-dialog>
</template>

<script setup>
import {onMounted, ref} from "vue";
import {TasksPage} from "../../../wailsjs/go/main/App.js";
import GoodsEdit from "./GoodsEditDialog.vue";
import {ElMessageBox} from "element-plus";
import {Delete, Edit} from "@element-plus/icons-vue";
import GoodsDeleteDialog from "./GoodsDeleteDialog.vue";
import {errElMessage} from "../../utils/el-message-utils.js";

//对话框可显示
const dialogVisible = ref(false)
const deleteDialogVisible = ref(false)
//表格分页响应结果
const tablePageVOList = ref([])
const pageTotal = ref()
const goodsTypeOpts = ref([])
const noAccept = ref(true)

//表格请求
const tasksPageRequest = ref({
  taskName: '',
  tasksStatus: 0,
  page: 1,
  size: 10
})

//props传参
const currentGoodsIdToDelete = ref(null)
const currentGoodsVO = ref(null)
//页面加载后执行
onMounted(async () => {
  try {
    await onSubmit()
    // goodsTypeOpts.value = await getGoodsTypeSelect()
  } catch (err) {
    console.error("err")
  }
})

// 查询按钮点击事件处理函数
async function onSubmit() {
  try {
    const resp = await TasksPage(tasksPageRequest.value);
    if (resp.data == null || resp.data.data == null) {
      tablePageVOList.value = []
    } else {
      tablePageVOList.value = resp.data.data;
      for (let i = 0; i < tablePageVOList.value.length; i++) {
        if (tablePageVOList.value[i].taskStatus === 0) {
          tablePageVOList.value[i].taskStatus = "未开始"
        } else if (tablePageVOList.value[i].taskStatus === 1) {
          tablePageVOList.value[i].taskStatus = "正在进行"
        } else {
          tablePageVOList.value[i].taskStatus = "已完成"
        }
      }
      pageTotal.value = resp.data.total
    }
  } catch (err) {
    errElMessage(err)
  }
}

// 添加分页切换事件处理函数（假设你已经实现了分页逻辑）
function handleCurrentChange(pageNumber) {
  tasksPageRequest.value.page = pageNumber
  onSubmit()
}

//编辑商品
function addGoods() {
  dialogVisible.value = true
  currentGoodsVO.value = null
}


//删除商品
function deleteGoods(id) {
  deleteDialogVisible.value = true
  currentGoodsIdToDelete.value = id
}

// 关闭删除对话框
function closeDeleteDialog() {
  deleteDialogVisible.value = false
}

//关闭编辑对话框
function closeDialog() {
  dialogVisible.value = false
}

//关闭处理
function handleClose() {
  ElMessageBox.confirm('是否关闭?')
      .then(() => {
        onSubmit()
      })
      .catch(() => {
        // catch error
      }).finally(() => {
    dialogVisible.value = false
    deleteDialogVisible.value = false
  })
}
</script>


<style scoped>
.form-select {
  --el-select-width: 180px;
}

.form-input {
  --el-select-width: 120px;
}
</style>