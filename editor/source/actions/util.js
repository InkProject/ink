import { ACTION } from './index';

function toogleTip(show, content) {
    return {
        type: ACTION.SHOW_TIP,
        show,
        content
    }
}

export function showTip(content) {
    return dispatch => {
        dispatch(toogleTip(true, content));
        setTimeout(function() {
            dispatch(toogleTip(false));
        }, 1000);
    };
}
