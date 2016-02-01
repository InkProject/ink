import * as util from '../util'
import * as listAction from '../list/action'
import ActionType from '../action'
import history from '../../store/history'

export function createArticle(name) {
    return dispatch => {
        dispatch(util.showTip('load'))
        util.apiRequest('POST', `articles`, {
            name,
            content: `
title: "未命名标题"
date: 2015-03-01 18:00:00 +0800

---

            `
        }).then(function(data) {
            dispatch(util.showTip('auto', '创建成功'))
            dispatch(listAction.fetchList())
            history.replaceState(null, `/edit/${data.id}`)
        })
    }
}
