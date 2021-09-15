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
                // badge: {
                //     color: 'primary',
                //     text: 'NEW'
                // }
            },
            {
                _name: 'CSidebarNavItem',
                name: i18n.t('message.products'),
                to: {name: 'ProductList'},
                icon: 'cil-applications',
            },
            // {
            //     _name: 'CSidebarNavItem',
            //     // name: 'Create Blog Post',
            //     name: i18n.t('message.create_blog_post'),
            //     to: {name: 'CreatePost'},
            //     icon: 'cil-image-plus',
            // },
        ]
    }
]
