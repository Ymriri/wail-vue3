<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="任务名称">
      <el-input v-model="form.taskName"/>
    </el-form-item>
    <el-form-item label="任务描述">
      <el-input v-model="form.taskDescription"/>
    </el-form-item>
    <el-form-item label="ftp访问路径">
      <el-input v-model="form.accessPath"/>
    </el-form-item>
    <el-form-item label="任务开始时间">
      <el-date-picker
          v-model="form.taskStartTime"
          type="datetime"
          placeholder="选择时间"
          :shortcuts="shortcuts"
      />
    </el-form-item>
    <el-form-item label="截止时间">
      <el-date-picker
          v-model="form.taskDeadline"
          type="datetime"
          placeholder="选择时间"
          :shortcuts="shortcuts"
      />
    </el-form-item>
    <el-form-item label="任务终止时间">
      <el-date-picker
          v-model="form.taskEndTime"
          type="datetime"
          placeholder="选择时间"
          :shortcuts="shortcuts"
      />
    </el-form-item>
    <el-form-item label="任务状态">
      <el-select v-model="form.taskStatus" placeholder="请选择任务状态" class="goodsType-select">
        <el-option v-for="item in taskTypeOpts" :key="item.id" :label="item.name" :value="item.id"/>
      </el-select>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">确认</el-button>
      <el-button @click="closeTaskDialog">取消</el-button>
    </el-form-item>
  </el-form>

</template>

<script setup>
import {defineEmits, onMounted, reactive, ref, watchEffect} from 'vue'
import {TasksSave, TaskUpdate} from '../../../wailsjs/go/main/App.js'
import {ElMessage} from "element-plus";
// import GoodsTypeDialog from "../goodstype/GoodsTypeDialog.vue";
import {commonElMessage, errElMessage, msgElMessage} from '../../utils/el-message-utils.js'

const emit = defineEmits(['refresh', 'close']);
//props接受参数
const props = defineProps({
  goodsVO: {
    type: [Object, null],
    required: true
  }
})
//商品类型对话框
// const goodsTypeDialogVisible = ref(false)
//判断是否为编辑
const isEdit = ref(false)
// 表单
let form = ref() // 初始化为空ref对象
const taskTypeOpts = ref([{"name": "未开始", "id": 0}, {"name": "正在进行", "id": 1}, {"name": "已完成", "id": 2}])

//页面加载后
onMounted(async () => {
  // 类型固定只有三种
  // await goodsTypes()

})

// 表单参数，这里使用watchEffect来监听props的变化
watchEffect(() => {
  if (props.goodsVO) {
    form.value = reactive(props.goodsVO)
    isEdit.value = true
    // 修改状态 为0-2

    console.log("开始编辑")
  } else {
    form.value = {
      taskName: '',
      taskDescription: '',
      taskStartTime: '',
      taskEndTime: '',
      taskDeadline: '',
      taskStatus: 0,
      accessPath: '',
    }
    isEdit.value = false
  }
})

// 取消
const closeTaskDialog = () => {
  emit('close')
}
// 查询列表
const onSubmit = () => {
  let taskSaveRequest = reactive({...form.value})
  if (isEdit.value) {
    let taskSaveRequest = {
      id: form.value.id,
      ...form.value
    }
    console.log(taskSaveRequest)
    if (0<=taskSaveRequest.taskStatus&&taskSaveRequest.taskStatus<=2) {
      taskSaveRequest.taskStatus = taskSaveRequest.taskStatus
    } else if (taskSaveRequest.taskStatus == '未开始') {
      taskSaveRequest.taskStatus = 0
    } else if (taskSaveRequest.taskStatus == '正在进行') {
      taskSaveRequest.taskStatus = 1
    } else{
      taskSaveRequest.taskStatus = 2
    }
    TaskUpdate(taskSaveRequest).then((resp) => {
      // 现实默认的消息
      commonElMessage(resp)
      // 这里应该取消弹窗的
      emit('close')
      //刷新分页
      emit('refresh')
      // console.log(resp)
    }).catch(err => {
      errElMessage(err)
    })
  } else {

    if (taskSaveRequest.taskName === '' || taskSaveRequest.taskDescription === '' || taskSaveRequest.taskStartTime === '' || taskSaveRequest.taskEndTime === '' || taskSaveRequest.taskDeadline === '' || taskSaveRequest.accessPath === '') {
      ElMessage.error('请填写完整信息')
      return
    }
    TasksSave(taskSaveRequest).then((resp) => {
      msgElMessage(resp, "添加任务成功！")
      emit('close')
      //刷新分页
      emit('refresh')
    }).catch(err => {
      ElMessage.error('出现异常' + err)
    })
  }
}


</script>


<style scoped>
.goodsType {
  display: flex;
}

.goodsType-select {
  --el-select-width: 220px;
}

.goodsType-button {
  margin-left: 1px;
}
</style>