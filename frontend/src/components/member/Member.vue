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
      <el-button type="warning" @click="openLoadConfigDialog">导入数据</el-button>
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

  <!--管理分组-->
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
  <!--  导入数据-->
  <el-dialog v-model="loadConfigVisible" title="数据导入" width="800">
    <div class="show-mid">
      <el-select
          v-model="tempAddConfigSelect"
          placeholder="请选择分组"
          clearable
          class="form-select"
          style="width: 120px"
      >
        <el-option v-for="item in ClassConfigSelect" :key="item.ID" :label="item.ConfigName" :value="item.ID"/>
      </el-select>
      <!--    添加 按钮，选择配置文件-->
      <!--  居中显示-->

      <el-button type="primary" @click="loadUserFile(null)" style="margin-left: 3%;">选择用户信息文件</el-button>
    </div>
    <el-divider></el-divider>
    <el-table :data="tempUser" v-loading="tempUserloading">
      <el-table-column prop="tempId" label="序号"/>
      <el-table-column prop="name" label="姓名"/>
      <el-table-column prop="employeeNumber" label="工号"/>
      <el-table-column prop="class" label="班级"/>
      <el-table-column prop="grade" label="年级"/>
      <el-table-column prop="department" label="部门"/>
      <el-table-column prop="configName" label="科组"/>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button v-if="scope.ID!==-1" type="danger" :icon="Delete" circle
                     @click="deleteLocalClassConfig(scope.row)"/>
        </template>
      </el-table-column>
    </el-table>
    <div class="show-mid" style="margin-top: 3%">
      <el-button type="primary" @click="addNewUserConfig(null)">导入</el-button>
    </div>
  </el-dialog>
</template>

<script setup>
import {onMounted, ref} from "vue";
import {
  ConfigDelete,
  ConfigInsert,
  MemberPage,
  SelectFile,
  UserBatchInsert,
  UserPage
} from "../../../wailsjs/go/main/App.js";
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
const loadConfigVisible = ref(false)


// 临时信息
const tempAddConfigSelect = ref([])
const tempUser = ref([])

//页面加载后执行
onMounted(async () => {
  reloaddd()
  searchPageUser()
})


async function reloaddd() {
  ClassConfigSelect.value = await getClassConfigSelect()
  ClassConfigSelect.value.push({ID: -1, ConfigName: '全部'})
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
    pageTotal.value = resp.data.total
    console.log(userPageVO.value)
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

function openLoadConfigDialog() {
  loadConfigVisible.value = true
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

const tempUserloading = ref(false)

async function loadUserFile() {
  tempUserloading.value = true
  let data = await SelectFile()
  tempUserloading.value = false
  for (let i = 0; i < data.length; i++) {
    data[i].tempId = i + 1
  }
  tempUser.value = data

}


function addNewUserConfig() {
  // 检查select 选择
  if (tempAddConfigSelect.value === '' || tempAddConfigSelect.value === -1) {
    ElMessage({
      message: '请选择正确的分组',
      type: 'warning',
      plain: true,
    })
    return
  }
  // 检查数据数量大于>1
  if (tempUser.value.length === 0) {
    ElMessage({
      message: '请选择正确的文件',
      type: 'warning',
      plain: true,
    })
    return
  }
  ElMessageBox.confirm(
      '是否确定导入数据？',
      '导入用户信息确定',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
  )
      .then(() => {

        for (let i = 0; i < tempUser.value.length; i++) {
          tempUser.value[i].configFileID = tempAddConfigSelect.value
        }
        UserBatchInsert(tempUser.value).then(resp => {
          // 刷新本地
          reloaddd()
          // 提醒删除成功
          ElMessage({
            message: '导入成功！',
            type: 'success',
            plain: true,
          })
          // 取消对话
          loadConfigVisible.value = false
          // 清空原来的数据
          tempUser.value = []
        })
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '取消导入',
        })
      })

}

function deleteLocalClassConfig(data) {
  ElMessageBox.confirm(
      '是否确定删除该用户？',
      '警告',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }
  )
      .then(() => {
        // 提醒删除成功
        for (let i = 0; i < tempUser.value.length; i++) {
          if (tempUser.value[i].tempId === data.tempId) {
            tempUser.value.splice(i, 1)
            break
          }
        }
        // 重新编号
        for (let i = 0; i < tempUser.value.length; i++) {
          tempUser.value[i].tempId = i + 1
        }
        ElMessage({
          message: '删除配置文件用户成功！',
          type: 'success',
          plain: true,
        })
      })
      .catch(() => {
        ElMessage({
          type: 'info',
          message: '取消删除',
        })
      })

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
  userPageRequest.value.page = pageNumber
  searchPageUser()
}

//打开删除对话框
function openDeleteDialog(row) {
  memberDeleteDialogVisible.value = true

  memberVO.value = row
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

.show-mid {
  display: flex;
  justify-content: center;
}
</style>