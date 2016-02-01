import * as listAction from '../list/action'
import * as util from '../util'
import history from '../../store/history'

export function saveContent() {
    return dispatch => {
        let state = globalStore.getState()
        let currentId = state.editor.get('id')
        let currentContent = state.editor.get('current')
        dispatch(util.showTip('load'))
        util.apiRequest('PUT', `articles/${currentId}`, {
            content: currentContent
        }).then(function(data) {
            dispatch(listAction.fetchList())
            dispatch(util.showTip('auto', '保存成功'))
        })
    }
}

export function removeArticle() {
    return dispatch => {
        let state = globalStore.getState()
        let currentId = state.editor.get('id')
        dispatch(util.showTip('load'))
        util.apiRequest('DELETE', `articles/${currentId}`).then(function(data) {
            history.replaceState(null, `/`)
            dispatch(listAction.fetchList())
            dispatch(util.showTip('auto', '删除成功'))
        })
    }
}
