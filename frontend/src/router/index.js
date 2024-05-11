import {createRouter, createWebHistory} from "vue-router";
import Task from "../components/task/Task.vue";
import Member from "../components/member/Member.vue";
import Bill from "../components/order/Order.vue";
import OrderRecord from "../components/order/OrderRecord.vue";
import System from "../components/system/System.vue";
import TaskDetail from "../components/taskdetail/TaskDetail.vue";

const routes = [
    {
        path: '/',
        name: 'Task',
        component: Task,
    }, {
        path: '/member',
        name: 'Member',
        component: Member,
    }, {
        path: '/order',
        name: 'Bill',
        component: Bill,
    }, {
        path: '/orderRecord',
        name: 'OrderRecord',
        component: OrderRecord,
    }, {
        path: '/system',
        name: 'System',
        component: System,
    }, {
        path: '/taskDetail',
        name: 'TaskDetail',
        component: TaskDetail,
    },
]
const router = createRouter({
    history: createWebHistory(), // 路由模式：history模式
    routes: routes
})
export default router