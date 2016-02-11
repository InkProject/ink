import * as util from '../util'
import ActionType from '../action'
import * as editorAction from '../editor/action'

export function showModal(flag) {
    return {
        type: ActionType.MODAL_SHOW,
        flag
    }
}

export function changeFocusMode(flag) {
    return {
        type: ActionType.MENU_CHNAGE_FOCUS_MODE,
        flag
    }
}

export function openConfig() {
    return dispatch => {
        (async () => {
            const data = await util.apiRequest('GET', `config`)
            dispatch(editorAction.setEditor('config', data))
        })()
    }
}

export function openHelp() {
    return dispatch => {
        (async () => {
            const data = ''
            dispatch(editorAction.setEditor('help', data))
        })()
    }
}
