import request from '../utils/request'
import ResponseData from '../utils/request'


export const deleteAppAPI = (app_id: number): Promise<typeof ResponseData> => {
    return request.delete('/app', { data: { app_id } })
}
