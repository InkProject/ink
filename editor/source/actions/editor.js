import { ACTION } from './index';

export function setContent(content) {
    return {
        type: ACTION.SET_CONTENT,
        content
    }
}
