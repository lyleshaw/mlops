import request from '../utils/request'
import ResponseData from '../utils/request'

export const loginAPI = (email: string, password: string): Promise<typeof ResponseData> => {
    return request.post('/user/login', { email, password })
}
