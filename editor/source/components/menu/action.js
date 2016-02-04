import * as util from '../util'
import ActionType from '../action'

export function showModal(flag) {
    return {
        type: ActionType.MODAL_SHOW,
        flag
    }
}
