import React from 'react'
import Component from './index'
import classNames from 'classnames'

import { connect } from 'react-redux'

class Toolbar extends Component {
    render() {
        const tip = this.props.util.get('tip')
        return (
            <div id="tooltip" className={classNames({hide: !tip.get('show'), error: tip.get('error')})}>
                <div className="content">
                    {tip.get('loading') ? <i className="fa fa-cog fa-spin"></i> : null}
                    {tip.get('content')}
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
