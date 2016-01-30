import { ACTION, utilAction, listAction } from './index'
import history from '../history'

export function createArticle(name) {
    return dispatch => {
        dispatch(utilAction.showTip('load'))
        utilAction.apiRequest('POST', `articles`, {
            name,
            content: `
title: "未命名标题"
date: 2015-03-01 18:00:00 +0800

---

            `
        }).then(function(data) {
            dispatch(utilAction.showTip('auto', '创建成功'))
            dispatch(listAction.fetchList())
            history.replaceState(null, `/edit/${data.id}`)
        })
    }
}
