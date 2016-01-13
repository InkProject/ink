import { ACTION, apiURL } from './index';

function toogleTip(show, loading, content) {
    return {
        type: ACTION.SHOW_TIP,
        loading,
        show,
        content
    }
}

export function showTip(type, content) {
    return dispatch => {
        if (type === 'auto') {
            dispatch(toogleTip(true, false, content));
            setTimeout(function() {
                dispatch(toogleTip(false));
            }, 1000);
        } else if (type === 'load') {
            dispatch(toogleTip(true, true, content));
        } else if (type === 'hide') {
            dispatch(toogleTip(false));
        }
    };
}

export function apiRequest(method, url, data) {
    let state = globalStore.getState();
    let currentId = state.editor.get('id');
    let currentContent = state.editor.get('current');
    return fetch(`${apiURL}/${url}`, {
        mode: 'cors',
        method: method,
        body: data ? JSON.stringify(data) : null
    }).then(function(response) {
        return response.json();
    }).catch(function(error) {
        showTip('load', error);
    });
}
