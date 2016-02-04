import * as listAction from './action'
import * as editorAction from '../editor/action'
import * as util from '../util'
import ActionType from '../action'

export function show() {
    return {
        type: ActionType.LIST_SHOW
    }
}

export function hide() {
    return {
        type: ActionType.LIST_HIDE
    }
}

function refresh(data) {
    return {
        type: ActionType.LIST_REFRESH,
        data
    }
}

export function fetch() {
    return dispatch => {
        (async () => {
            const data = await util.apiRequest('GET', `articles`)
            dispatch(refresh(data))
        })()
    }
}

export function open(id) {
    return dispatch => {
        (async () => {
            const data = await util.apiRequest('GET', `articles/${id}`)
            dispatch(editorAction.setEditor(id, data))
        })()
    }
}
