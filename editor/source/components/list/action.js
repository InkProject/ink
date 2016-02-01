import * as listAction from './action'
import * as editorAction from '../editor/action'
import * as util from '../util'
import ActionType from '../action'

export function showLoading(flag) {
    return {
        type: ActionType.SHOW_LOADING,
        flag
    }
}

function refreshList(data) {
    return {
        type: ActionType.REFRESH_LIST,
        data
    }
}

export function fetchList() {
    return dispatch => {
        // dispatch(util.showTip('load'))
        util.apiRequest('GET', `articles`).then(function(data) {
            dispatch(refreshList(data))
            // dispatch(util.showTip('hide'))
        })
    }
}

export function showList() {
    return {
        type: ActionType.SHOW_LIST
    }
}

export function hideList() {
    return {
        type: ActionType.HIDE_LIST
    }
}

export function openArticle(id) {
    return dispatch => {
        dispatch(util.showTip('load'))
        util.apiRequest('GET', `articles/${id}`).then(function(data) {
            dispatch(editorAction.setEditor(id, data))
            dispatch(util.showTip('hide'))
        })
    }
}
