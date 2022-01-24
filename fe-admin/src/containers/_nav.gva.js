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
                exact: false, //[extract=false] to make this item "active" on child-action (create/edit/delete)
            },
            {
                _name: 'CSidebarNavItem',
                name: i18n.t('message.users'),
                to: {name: 'UserList'},
                icon: 'cil-group',
                exact: false, //[extract=false] to make this item "active" on child-action (create/edit/delete)
            },
        ]
    }
]
