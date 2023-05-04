import request from '../utils/request'
import ResponseData from '../utils/request'


export const getAppAPI = (): Promise<typeof ResponseData> => {
    return request.get('/app/me')
}
