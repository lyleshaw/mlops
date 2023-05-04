import request from '../utils/request'
import ResponseData from '../utils/request'


export const sendVerificationCodeAPI = (email: string): Promise<typeof ResponseData> => {
    return request.post('/email/code', { email })
}
