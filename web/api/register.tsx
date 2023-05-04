import request from '../utils/request'
import ResponseData from '../utils/request'


export const registerAPI = (username: string, email: string, password: string, code: string): Promise<typeof ResponseData> => {
    return request.post('/user/register', { username, email, password, code })
}
