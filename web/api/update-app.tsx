import request from '../utils/request'
import ResponseData from '../utils/request'


export const updateAppAPI = (app_id: number, app_name: string, app_image: string, is_reload: boolean): Promise<typeof ResponseData> => {
    return request.patch('/app', { app_id, app_name, app_image, is_reload })
}
