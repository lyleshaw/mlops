import request from '../utils/request'
import ResponseData from '../utils/request'


export const createAppAPI = (app_name: string, app_image: string): Promise<typeof ResponseData> => {
    return request.post('/app', { app_name, app_image })
}
