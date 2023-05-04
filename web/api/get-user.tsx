import request from '../utils/request'
import ResponseData from '../utils/request'


export const getUserAPI = (): Promise<typeof ResponseData> => {
    return request.get('/user/me')
}
