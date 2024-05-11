<template>
  <h3>任务详情</h3>
  <el-container>
    <el-header height="370px">
      <el-form :model="nowTask" class="form" label-position="left">
        <!--        设置在同一行-->
        <div style="display:flex;justify-content: space-between;">
          <el-form-item label="任务ID：" label-width="110px">
            <el-input v-model="nowTask.id" disabled/>
          </el-form-item>
          <el-form-item label="ftp访问路径：" label-width="110px">
            <el-input v-model="nowTask.accessPath" disabled/>
          </el-form-item>
          <el-form-item label="选择分组：" label-width="110px">
            <el-select
                v-model="newClassConfig.configFileID"
                placeholder="请选择分组"
                clearable
                class="form-select"
                style="width: 196px"
            >
              <el-option v-for="item in ClassConfigSelect" :key="item.ID" :label="item.ConfigName" :value="item.ID"/>
            </el-select>
          </el-form-item>
        </div>
        <el-form-item label="任务名称：" label-width="110px">
          <el-input v-model="nowTask.taskName" disabled/>
        </el-form-item>
        <el-form-item label="任务描述：" label-width="110px">
          <el-input v-model="nowTask.taskDescription" disabled/>
        </el-form-item>
        <!--        设置在同一行-->
        <div style="display: flex;justify-content: space-between;">
          <el-form-item label="开始时间：" label-width="110px">
            <el-input v-model="nowTask.taskStartTime" disabled/>
          </el-form-item>
          <el-form-item label="截止日期：" label-width="110px">
            <el-input v-model="nowTask.taskDeadline" disabled/>
          </el-form-item>
          <el-form-item label="任务终止时间：" label-width="110px">
            <el-input v-model="nowTask.taskEndTime" disabled/>
          </el-form-item>
        </div>
        <!--        设置在同一行对齐-->

        <div style="display: flex;margin: auto;">


        </div>
        <div style="display: flex;margin: auto;">
          <el-form-item label="文件格式：" label-width="110px">
            <el-input v-model="nowTask.accessPath"/>
          </el-form-item>
          <!--          <el-form-item style="margin-left: 1%">-->
          <!--            <el-input v-model="nowTask.month" style="width: 10px"/>-->
          <!--            <el-text>月</el-text>-->
          <!--          </el-form-item>-->
          <!--          <el-form-item label="" label-width="10px">-->
          <!--            <el-input v-model="nowTask.count"/>-->
          <!--          </el-form-item>-->
          <!--          <el-form-item label="周" label-width="10px">-->
          <!--            <el-input v-model="nowTask.week"/>-->
          <!--          </el-form-item>-->
          <el-button type="primary" style="margin-left: 7%">预览生成</el-button>
          <el-button type="primary" style="margin-left: 7%">确定格式</el-button>
          <el-button type="primary" style="margin-left: 7%">生成子任务</el-button>
        </div>
        <div style="display: flex;margin: auto;">
          <el-form-item label="任务状态：" label-width="110px">
            <el-input v-model="nowTask.taskLable" disabled/>
          </el-form-item>
          <el-button :type="nowTaskButtonTitle.type" style="margin-left: 7%">{{ nowTaskButtonTitle.title }}
          </el-button>
        </div>
        <el-divider/>
        <!--        路径生成-->
      </el-form>
    </el-header>

    <el-main>
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span>子任务详情</span><br/><br/>
            <!--            横向-->
            <div style="display: flex; flex-direction: row">
              <el-text style="width: 120px">子任务状态：</el-text>
              <el-select
                  v-model="taskDetailPage.status"
                  placeholder="请选择子任务状态"
                  clearable
                  class="form-select"
                  style="width: 196px"
              >
                <el-option v-for="item in filterTableSearch" :key="item.id" :label="item.name" :value="item.id"/>
              </el-select>
              <div class="table-search">
                <el-input v-model="taskDetailPage.count" style="width: 50px"/>
                <el-text style="margin-left: 1%">次</el-text>
              </div>
              <div class="table-search">
                <el-input v-model="taskDetailPage.month" style="width: 50px"/>
                <el-text style="margin-left: 1%">月</el-text>
              </div>
              <div class="table-search">
                <el-input v-model="taskDetailPage.week" style="width: 50px;"/>
                <el-text style="margin-left: 2%">周</el-text>
              </div>
              <el-button type="primary" style="margin-left: 7.8%">查询任务</el-button>
              <el-button type="warning" style="margin-left: 7.2%">导出文件</el-button>
            </div>
          </div>

        </template>
        <div>
          <el-table :data="bills" stripe style="width: 100%">
            <el-table-column prop="goodsName" label="子任务ID"/>
            <el-table-column prop="name" label="姓名"/>
            <el-table-column prop="employeeNumber" label="工号"/>
            <el-table-column prop="class" label="班级"/>
            <el-table-column prop="grade" label="年级"/>
            <el-table-column prop="department" label="部门"/>
            <el-table-column prop="configName" label="科组"/>
            <el-table-column prop="configName" label="格式对比"/>
            <el-table-column prop="configName" label="文件名"/>
            <el-table-column prop="configName" label="状态"/>
          </el-table>
        </div>
      </el-card>
    </el-main>
  </el-container>

  <!--  用户余额不足对话框  -->
  <el-dialog
      v-model="dialogVisible"
      title="余额不足"
      width="500"
  >
    <span>用户余额不足,是否直接抵扣(完全扣除余额)，用户余额{{ memberForm.account }}元</span>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="memberDeduction">
          确认
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>

import {onMounted, onUnmounted, ref} from "vue";
import {allMemberSelect, getClassConfigSelect} from "../../api/select.js";
import {GoodsGetById, MemberGetById, Order, TasksGetById} from "../../../wailsjs/go/main/App.js";
import {ElMessage} from "element-plus";
import {errElMessage, msgElMessage} from "../../utils/el-message-utils.js";

import {useRoute} from 'vue-router';

// 传过来的数据
let nowTaskId = ref(-1);

let nowTaskButtonTitle = ref({
  "title": "开始任务",
  "type": "primary"
})

//对话框
const dialogVisible = ref(false)
//是否抵扣
const isDeduction = ref(false)
const ClassConfigSelect = ref([])

// 分组 select
const newClassConfig = ref({
  ConfigName: '',
  ParentID: 3,
})

const filterTableSearch = ref([{
  name: "所有",
  id: -1
}, {
  name: '已提交',
  id: 0
}, {
  name: '迟交',
  id: 1
}, {
  name: "异常文件",
  id: 2
}
])

const taskDetailPage = ref({
  taskName: '',
  status: '',
  month: '',
  count: '',
  week: ''
})

const form = ref({
  memberId: '',
  realPrice: '',
  num: 1,
})
const goodsForm = ref({
  id: '',
  name: '',
  goodsNumber: '',
  goodsType: '',
  price: '',
  count: ''
})
const memberForm = ref({
  id: '',
  name: '',
  phone: '',
  level: '',
  account: '',
  discount: ''
})
const bills = ref([])
const totalPrice = ref('0')
const pay = ref('0')
const goodsSelect = ref([])
const memberSelect = ref([])
// 当前任务
const nowTask = ref({
  id: '',
  taskName: '',
  taskDescription: '',
  accessPath: '',
  startTime: '',
  endTime: '',
  taskStatus: '',
  taskLable: '',
  count: '',
  week: '',
  month: '',
});

onMounted(async () => {
  // console.log("加载...")
  // 只会进来一次
  if (useRoute().query.id) {
    // nowTaskId.value =
    // console.log(nowTaskId.value)
    nowTaskId.value = useRoute().query.id
    getNowTaskMessage(useRoute().query.id)
  }
  ClassConfigSelect.value = await getClassConfigSelect()
})

// 需要接触的参数
onUnmounted(async () => {
  // 是否需要撤销呢
  // nowTaskId.value=-1
  console.log("页面消失")
})


function getNowTaskMessage(id) {
  // 转成int64类型
  id = parseInt(id)
  TasksGetById(id).then(resp => {
    nowTask.value = resp.data
    if (nowTask.value.taskStatus === 0) {
      nowTask.value.taskLable = "未开始"
      nowTaskButtonTitle.value.title = "开始任务"
      nowTaskButtonTitle.value.type = "primary"
    } else if (nowTask.value.taskStatus === 1) {
      nowTask.value.taskLable = "正在进行"
      nowTaskButtonTitle.value.title = "结束任务"
      nowTaskButtonTitle.value.type = "danger"
    } else {
      nowTask.value.taskLable = "已完成"
      nowTaskButtonTitle.value.title = "重新开始"
      nowTaskButtonTitle.value.type = "warning"
    }
    console.log(nowTask.value)
  }).catch(err => {
    errElMessage(err)
  })
}


// 跳转页面加载参数

// 下单
function onSubmit() {
  let request = {
    pay: totalPrice.value,
    memberId: memberForm.value.id,
    orderItem: []
  }
  for (let i = 0; i < bills.value.length; i++) {
    let bill = bills.value[i]
    let order = {
      goodsId: bill.goodsId,
      realPrice: bill.realPrice,
      num: bill.num
    }
    request.orderItem.push(order)
  }
  //如果购物车价格超过会员账户，提示
  if (!dialogVisible.value && memberForm.value.id !== "" && Number(request.pay) > Number(memberForm.value.account)) {
    dialogVisible.value = true
  } else {
    //下订单
    Order(request).then(resp => {
      msgElMessage(resp, "下单成功")
      bills.value = []
      totalPrice.value = 0
      isDeduction.value = false
    }).catch(err => {
      errElMessage(err)
    })
  }
}

//加入购物车
function addCart() {
  let bill = {
    goodsId: goodsForm.value.id,
    goodsName: goodsForm.value.name,
    goodsNumber: goodsForm.value.goodsNumber,
    goodsType: goodsForm.value.goodsType,
    price: goodsForm.value.price,
    realPrice: form.value.realPrice,
    num: form.value.num,
  }
  if (goodsForm.value.count < bill.num) {
    ElMessage.warning(goodsForm.value.name + "库存不足，库存数量为" + goodsForm.value.count)
    return
  }
  for (let i = 0; i < bills.value.length; i++) {
    if (bills.value[i].goodsId === bill.goodsId) {
      ElMessage.warning(bills.value[i].goodsName + "商品已加入购物车")
      return
    }
  }
  bills.value.push(bill)
  calcTotalPrice()
}

// 切换商品
function changeGoods(goodsId) {
  GoodsGetById(goodsId).then(resp => {
    goodsForm.value.id = resp.data.id
    goodsForm.value.name = resp.data.name
    goodsForm.value.goodsType = resp.data.goodsType
    goodsForm.value.goodsNumber = resp.data.goodsNumber
    goodsForm.value.price = resp.data.price
    goodsForm.value.count = resp.data.count
    calcRealPrice()
  })
}

//deleteGoods 删除商品
function deleteGoods(row) {
  for (let i = 0; i < bills.value.length; i++) {
    if (bills.value[i].goodsId === row.goodsId) {
      bills.value.splice(i, 1)
      break
    }
  }
  calcTotalPrice()
}

// 根据手机号过滤
async function filterMember(phone) {
  memberSelect.value = await allMemberSelect()
  if (phone) {
    let select = []
    for (let i = 0; i < memberSelect.value.length; i++) {
      let member = memberSelect.value[i]
      if (member.phone.includes(phone)) {
        select.push(member)
      }
    }
    memberSelect.value = select
  }
}

// 切换会员
function changeMember(memberId) {
  MemberGetById(memberId).then(resp => {
    memberForm.value.id = resp.data.id
    memberForm.value.name = resp.data.name
    memberForm.value.level = resp.data.level.name
    memberForm.value.discount = resp.data.level.discount
    memberForm.value.account = resp.data.account
    calcRealPrice()
  })
}

//修改商品数量
function changeGoodsNum() {
  calcTotalPrice()
}

//计算实际支付
function calcRealPrice() {
  if (goodsForm.value.id !== "") {
    if (memberForm.value.id !== "" && memberForm.value.id !== '0') {
      form.value.realPrice = (Number(goodsForm.value.price) * Number(memberForm.value.discount)).toFixed(2)
    } else {
      form.value.realPrice = Number(goodsForm.value.price).toFixed(2)
    }
  }
}

//计算总价
function calcTotalPrice() {
  totalPrice.value = '0'
  for (let i = 0; i < bills.value.length; i++) {
    let bill = bills.value[i]
    totalPrice.value = (Number(totalPrice.value) + Number(bill.realPrice) * bill.num).toFixed(2)
  }
}

//会员抵扣对话框确认
function memberDeduction() {
  isDeduction.value = true
  onSubmit()
  dialogVisible.value = false
}
</script>


<style scoped>
.table-search {
  display: flex;
  margin-left: 3%;
}
</style>