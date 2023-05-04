import axios from 'axios';

// const BASE_URL = 'http://localhost:7654/api/v1';
const BASE_URL = 'https://mlops-api.aireview.tech/api/v1';

interface ResponseData {
    message: string;
    code: number;
    data: any;
}

// 创建axios实例
const service = axios.create({
    baseURL: BASE_URL, // api的base_url
    timeout: 60000, // 请求超时时间，1分钟
    withCredentials: true,
});

// request拦截器
service.interceptors.request.use(
    (config) => {
        // Do something before request is sent
        if (localStorage.getItem('token')) {
            config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`;
        }
        console.log(config);
        return config;
    },
    (error) => {
        Promise.reject(error);
    },
);

// respone拦截器
service.interceptors.response.use(
    (response) => {
        /**
         * 下面的注释为通过response自定义code来标示请求状态，当code返回如下情况为权限有问题，登出并返回到登录页
         * 如通过xmlhttprequest 状态码标识 逻辑可写在下面error中
         */

        // const res = response.data;
        // if (response.status === 401 || res.status === 40101) {
        //   MessageBox.confirm('你已被登出，可以取消继续留在该页面，或者重新登录', '确定登出', {
        //     confirmButtonText: '重新登录',
        //     cancelButtonText: '取消',
        //     cancelButtonClass: 'calBtnClass',
        //     type: 'warning',
        //   }).then(() => {
        //     store.dispatch('FedLogOut').then(() => {
        //       location.reload(); // 为了重新实例化vue-router对象 避免bug
        //     });
        //   });
        //   return Promise.reject('error');
        // }

        // if (res.status === 40001) {
        //   Message({ message: '账户或密码错误！', type: 'warning' });
        //   return Promise.reject('error');
        // }

        // if (response.status !== 200 && res.status !== 200) {
        //   //此处并不能拦截全部异常情况，因为res是后台接口返回的数据，有一些由于不规范，没有带status属性，导致没有被拦截
        //   Message({ message: res.message, type: 'error', duration: 6 * 1000 });
        // } else {
        //   //只有成功后，才可以进行操作
        //   return response.data;
        // }
        return response.data;
    },
    (error) => {
        // Message({
        //   message: error.message,
        //   type: 'error',
        //   duration: 6 * 1000,
        // });
        return Promise.reject(error);
    },
);

export default service;