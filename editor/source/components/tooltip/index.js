import React from 'react'
import Component from '../index'
import classNames from 'classnames'

import { connect } from 'react-redux'

class Toolbar extends Component {
    render() {
        const tip = this.props.util.get('tip').toJS()
        return (
            <div id="tooltip" className={classNames({hide: !tip.show, error: tip.error})}>
                <div className="content">
                    {tip.loading ? <i className="fa fa-cog fa-spin"></i> : null}
                    {tip.content}
                    {tip.action ? <a className="action" onClick={tip.action.callback}>{tip.action.button}</a> : null}
                </div>
            </div>
        )
    }
}

export default connect(function(state) {
    return {
        util: state.util
    }
})(Toolbar)
