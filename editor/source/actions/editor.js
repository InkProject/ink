import { ACTION } from './index';
import yaml from 'js-yaml';
import _ from 'lodash';

let parseConfig = function(data, noContent) {
    let configStr;
    if (noContent) {
        configStr = _.trim(data);
    } else {
        configStr = _.trim(data.split('---')[0] || '\n');
    }
    try {
        return yaml.safeLoad(configStr);
    } catch (err) {
        console.log(err);
        return null;
    }
};

export function setHeader(data) {
    let config = parseConfig(data);
    return {
        type: ACTION.SET_HEADER,
        title: config ? config.title : '键入文章标题',
        tags: config ? config.tags : []
    }
}

export function setEditor(id, data) {
    let dataAry = data.split('---');
    let configData = _.trim(dataAry[0]);
    let content = _.trim(dataAry[1] ? dataAry.slice(1).join('---') : '');
    let config = parseConfig(configData);
    let { title, tags } = config || {};
    return {
        type: ACTION.SET_CONTENT,
        id,
        title,
        tags,
        config: configData,
        content
    }
}

export function setCurrent(current) {
    return {
        type: ACTION.SET_CURRENT,
        current
    }
}
