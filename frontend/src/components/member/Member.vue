<template>
  <!--  查询表单  -->
  <el-form :inline="true" :model="userPageRequest">
    <el-form-item label="姓名">
      <el-input v-model="userPageRequest.name" placeholder="请输入姓名" class="form-input" clearable/>
    </el-form-item>
    <el-form-item label="用户分组">
      <el-select
          v-model="userPageRequest.configFileID"
          placeholder="请选择分组"
          clearable
          class="form-select"
          style="width: 120px"
      >
        <el-option v-for="item in ClassConfigSelect" :key="item.ID" :label="item.ConfigName" :value="item.ID"/>
      </el-select>
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="searchPageUser">查询</el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="openClassConfigDialog(null)">管理分组</el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="warning">导入数据</el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="success" @click="reloaddd">刷新</el-button>
    </el-form-item>
  </el-form>

  <!--查询表格-->
  <el-table :data="userPageVO" stripe style="width: 100%">
    <el-table-column prop="id" label="ID"/>
    <el-table-column prop="name" label="姓名"/>
    <el-table-column prop="employeeNumber" label="工号"/>
    <el-table-column prop="class" label="班级"/>
    <el-table-column prop="grade" label="年级"/>
    <el-table-column prop="department" label="部门"/>
    <el-table-column prop="configName" label="科组"/>
    <el-table-column prop="configFileId" label="ConfigID"/>
    <el-table-column label="操作">
      <template #default="scope">
        <el-button type="primary" :icon="Edit" circle @click="openEditDialog(scope.row)"/>
        <el-button type="danger" :icon="Delete" circle @click="openDeleteDialog(scope.row)"/>
      </template>
    </el-table-column>
  </el-table>
  <!-- 如果需要根据的实际长度来控制分页组件 -->
  <el-pagination
      layout="prev, pager, next"
      :total="pageTotal"
      @current-change="handleCurrentChange"
  />
  <!-- 会员编辑对话框 -->
  <el-dialog
      v-model="memberEditDialogVisible"
      title="会员编辑"
      :before-close="memberEditDialogClose"
  >
    <member-edit-dialog
        :member-v-o="memberVO"
        @close="closeEditDialog"
    />
  </el-dialog>
  <!-- 会员删除对话框 -->
  <el-dialog
      v-model="memberDeleteDialogVisible"
      title="删除会员"
  >
    <member-delete-dialog
        :member-v-o="memberVO"
        @close="closeDeleteDialog"
    />
  </el-dialog>

  <el-dialog v-model="classConfigVisible" title="管理分组" width="500">
    <el-table :data="ClassConfigSelect">
      <el-table-column property="ID" label="ID" width="150"/>
      <el-table-column property="ConfigName" label="分组名" width="200"/>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button v-if="scope.ID!==-1" type="danger" :icon="Delete" circle @click="deleteClassConfig(scope.row)"/>
        </template>
      </el-table-column>

    </el-table>
    <!--    添加分割线-->
    <el-divider></el-divider>
    <el-form :inline="true" :model="newClassConfig">
      <el-form-item label="新增分组">
        <el-input v-model="newClassConfig.ConfigName" placeholder="请输入分组名" class="form-input" clearable/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="addNewClassConfig(null)">新增</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup>
import {onMounted, ref} from "vue";
import {ConfigDelete, ConfigInsert, MemberPage,UserPage} from "../../../wailsjs/go/main/App.js";
import {ElMessage, ElMessageBox} from "element-plus";
import {Delete, Edit} from "@element-plus/icons-vue";
import MemberEditDialog from "./MemberEditDialog.vue";
import MemberDeleteDialog from "./MemberDeleteDialog.vue";
import {getClassConfigSelect} from "../../api/select.js";
//表格请求
const memberPageRequest = ref({
  name: '',
  phone: '',
  configId: '',
  account: '',
  page: 1,
  size: 10
})

const userPageRequest = ref({
  name: '', // 姓名
  configFileID: '', // 分组情况
  page: 1,
  size: 10
})

const newClassConfig = ref({
  ConfigName: '',
  ParentID: 3,
})

const memberPageVO = ref([])

const userPageVO = ref([])

const pageTotal = ref(0)
const memberVO = ref(null)
const memberLevelSelect = ref([])

const ClassConfigSelect = ref([])

//会员编辑对话框
const memberEditDialogVisible = ref(false)
const memberDeleteDialogVisible = ref(false)
const classConfigVisible = ref(false)

//页面加载后执行
onMounted(async () => {

  // searchPageMember()
  reloaddd()
  // console.log(ClassConfigSelect)
  // console.log(ClassConfigSelect.value)
})


async function reloaddd() {
  console.log("点击刷新~")
  // searchPageMember()
  ClassConfigSelect.value = await getClassConfigSelect()
  // console.log(ClassConfigSelect)
  ClassConfigSelect.value.push({ID: -1, ConfigName: '全部'})
  console.log(ClassConfigSelect)
}

//搜索会员
function searchPageMember() {
  MemberPage(memberPageRequest.value).then(resp => {
    memberPageVO.value = resp.data.data
  })
}

// search user
function searchPageUser() {
  // 整数
  if (userPageRequest.value.configFileID === '') {
    userPageRequest.value.configFileID = -1
  }
  userPageRequest.value.configFileID = parseInt(userPageRequest.value.configFileID)
  UserPage(userPageRequest.value).then(resp => {
    userPageVO.value = resp.data.data
    console.log(userPageVO.value )
  })

}
//打开编辑对话框
function openEditDialog(row) {
  memberEditDialogVisible.value = true
  if (row) {
    memberVO.value = row
  } else {
    memberVO.value = null
  }
}

function addNewClassConfig() {
  // 添加新的分组
  ConfigInsert(newClassConfig.value).then(resp => {
    // 刷新本地
    reloaddd()
    // 提醒添加成功
    ElMessage({
      message: '添加分组成功！',
      type: 'success',
      plain: true,
    })
    // 清空原来的
    newClassConfig.value.ConfigName = ''
  })
  // classConfigVisible.value=false
}

function deleteClassConfig(data) {
  // 删除分组
  ElMessageBox.confirm(
      '是否确定删除分组？',
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
  )
      .then(() => {
        ConfigDelete(data.ID).then(resp => {
          // 刷新本地
          reloaddd()
          // 提醒删除成功
          ElMessage({
            message: '删除分组成功！',
            type: 'success',
            plain: true,
          })
        })
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '取消删除',
        })
      })

}

function openClassConfigDialog() {
  classConfigVisible.value = true
}

// 添加分页切换事件处理函数（假设你已经实现了分页逻辑）
function handleCurrentChange(pageNumber) {
  memberPageRequest.value.page = pageNumber
  searchPageMember()
}

//删除编辑对话框
function closeEditDialog() {
  memberEditDialogVisible.value = false
  searchPageMember()
}

//打开删除对话框
function openDeleteDialog(row) {
  memberDeleteDialogVisible.value = true
  memberVO.value = row
}

//删除对话框
function closeDeleteDialog() {
  memberDeleteDialogVisible.value = false
  searchPageMember()
}

//关闭会员编辑对话框
function memberEditDialogClose() {
  ElMessageBox.confirm('是否关闭?')
      .then(() => {
        //重新搜索会员
        searchPageMember()
      })
      .catch(() => {
        // catch error
      }).finally(() => {
    memberEditDialogVisible.value = false
  })
}
</script>


<style scoped>
</style>