import { ACTION, apiURL } from './index';

function toogleTip({
    show = true,
    error = false,
    loading = false,
    content = ''
}) {
    return {
        type: ACTION.SHOW_TIP,
        loading,
        show,
        error,
        content
    }
}

export function showTip(type, content) {
    return dispatch => {
        if (type === 'auto') {
            dispatch(toogleTip({
                content
            }));
            setTimeout(function() {
                dispatch(toogleTip({
                    show: false
                }));
            }, 1000);
        } else if (type === 'load') {
            dispatch(toogleTip({
                loading: true,
                content
            }));
        } else if (type === 'hide') {
            dispatch(toogleTip({
                show: false
            }));
        } else if (type === 'error') {
            dispatch(toogleTip({
                error: true,
                content
            }));
            setTimeout(function() {
                dispatch(toogleTip({
                    show: false
                }));
            }, 1000);
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
        showTip('error', error.message || error || '未知错误');
    });
}
