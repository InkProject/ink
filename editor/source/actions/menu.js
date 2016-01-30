import { ACTION, utilAction } from './index'

export function showModal(flag) {
    return {
        type: ACTION.SHOW_MODAL,
        flag
    }
}
