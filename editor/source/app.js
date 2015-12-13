import 'font-awesome/less/font-awesome.less';
import './styles/global.less';
import './styles/index.less';

import React from 'react';
import ReactDom from 'react-dom';

import Editor from './components/editor';

ReactDom.render(
    <div id="container">
        <div id="left">
            <ul className="menu">
                <li className="button button-circle create"><i className="fa fa-plus"></i></li>
                <li className="button button-circle setting"><i className="fa fa-wrench"></i></li>
                <li className="button button-circle"><i className="fa fa-hashtag"></i></li>
            </ul>
            <div id="files">
                <label id="search-wrap" htmlFor="search">
                    <i className="fa fa-search"></i>
                    <input id="search" type="text" placeholder="搜索..." />
                </label>
                <ul id="list">
                    <li className="item">
                        <div className="head">
                            <span className="date">2015-03-01 18:00</span>
                        </div>
                        <div className="title">构建只为纯粹书写的博客</div>
                        <div className="preview">纸小墨（InkPaper）是一个使用GO语言编写的静态博客构建工具，可以快速搭建博客网站。优点是无依赖跨平台，配置简单构建快速，注重简洁易用与排版优化</div>
                    </li>
                    <li className="item">
                        <div className="head">
                            <span className="date">2015-03-01 18:00</span>
                        </div>
                        <div className="title">构建只为纯粹书写的博客</div>
                        <div className="preview">纸小墨（InkPaper）是一个使用GO语言编写的静态博客构建工具，可以快速搭建博客网站。优点是无依赖跨平台，配置简单构建快速，注重简洁易用与排版优化</div>
                    </li>
                    <li className="item">
                        <div className="head">
                            <span className="date">2015-03-01 18:00</span>
                        </div>
                        <div className="title">构建只为纯粹书写的博客</div>
                        <div className="preview">纸小墨（InkPaper）是一个使用GO语言编写的静态博客构建工具，可以快速搭建博客网站。优点是无依赖跨平台，配置简单构建快速，注重简洁易用与排版优化</div>
                    </li>
                </ul>
            </div>
        </div>
        <div id="right">
            <ul className="menu">
                <li className="button button-cube"><i className="fa fa-rocket"></i>发布</li>
                <li className="button button-cube deploy"><i className="fa fa-chrome"></i>预览</li>
                <li className="button button-circle"><i className="fa fa-floppy-o"></i></li>
                <li className="button button-circle remove"><i className="fa fa-trash"></i></li>
            </ul>
        </div>
        <div id="header">
            <div className="title">构建只为纯粹书写的博客 —— 纸小墨</div>
            <div className="info">
                <span className="draft">草稿</span>
                <span className="tag">产品</span>
                <span className="tag">设计</span>
            </div>
        </div>
        <Editor></Editor>
    </div>
, document.getElementById('root'));
