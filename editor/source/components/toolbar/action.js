import * as listAction from '../list/action'
import * as util from '../util'
import history from '../../store/history'

export function saveContent() {
    return dispatch => {
        let state = store.getState()
        let currentId = state.editor.get('id')
        let currentContent = state.editor.get('current');
        (async () => {
            const data = await util.apiRequest('PUT', `articles/${currentId}`, { content: currentContent })
            dispatch(listAction.fetch())
            dispatch(util.showTip('auto', '保存成功'))
        })()
    }
}

export function removeArticle() {
    return dispatch => {
        let state = store.getState()
        let currentId = state.editor.get('id');
        (async () => {
            await util.apiRequest('DELETE', `articles/${currentId}`)
            dispatch(listAction.fetch())
            dispatch(util.showTip('auto', '删除成功'))
            history.push(`/`)
        })()
    }
}
