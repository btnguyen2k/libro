//#GovueAdmin-Customized
import i18n from '../i18n'

export default [
    {
        _name: 'CSidebarNav',
        _children: [
            {
                _name: 'CSidebarNavItem',
                name: i18n.t('message.dashboard'),
                to: {name: 'Dashboard'},
                icon: 'cil-wallpaper',
            },
            {
                _name: 'CSidebarNavItem',
                name: i18n.t('message.products'),
                to: {name: 'ProductList'},
                icon: 'cil-applications',
                exact: false, //set extract=false to this item "active" for sub-actions (create/edit/delete)
            },
        ]
    }
]
