//#GovueAdmin-Customized
import Vue from 'vue'
import VueI18n from 'vue-i18n'

const messages = {
    en: {
        _name: 'English',
        _flag: 'cif-gb',
        _demo_msg: 'This is instance is for demo purpose only. Data might be reset without notice.' +
            '<br/>Login with default account <strong>admin@local/s3cr3t</strong> ' +
            'or using your <u>social account</u> via "Login with social account" link ' +
            '(the application <u>will not</u> know or store your social account credential).' +
            '<br/><br/>You can also access the frontend via <a href="/doc/" style="color: yellowgreen">this link</a>.',

        message: {
            actions: 'Actions',
            action_create: 'Create',
            action_save: 'Save',
            action_back: 'Back',
            action_edit: 'Edit',
            action_delete: 'Delete',
            action_move_down: 'Move down',
            action_move_up: 'Move up',

            error: 'Error',
            info: 'Information',
            settings: 'Settings',
            account: 'Account',
            language: 'Language',

            error_field_mandatory: 'Field value is mandatory.',

            yes: 'Yes',
            no: 'No',
            ok: 'Ok',
            cancel: 'Cancel',
            close: 'Close',

            logout: 'Logout',
            login: 'Login',
            login_info: 'Please sign in to continue',
            login_social: 'Login with social account',
            wait: 'Please wait...',
            error_parse_login_token: 'Error parsing login-token',
            home: 'Home',
            dashboard: 'Dashboard',

            icons: "Icons",
            icon_icon: "Icon",
            icon_name: "Name",

            products: "Products",
            topics: "Topics",
            pages: "Pages",

            product_info: "Information",
            product_id: "Id",
            product_id_msg: "Must be unique, will be generated if empty",
            product_is_published: "Published",
            product_is_published_msg: "Product's documents are visible only when published",
            product_name: "Name",
            product_name_msg: "Display name of the product",
            product_desc: "Description",
            product_desc_msg: "Summary description of the product",
            product_domains: "Domain names",
            product_domains_msg: "Product documents are accessible via these domain names (one domain per line)",

            product_contacts: "Contacts",
            product_website: "Website",
            product_website_msg: "Product website",
            product_email: "Email",
            product_email_msg: "Contact email address",
            product_github: "Github",
            product_github_msg: "Url of GitHub page",
            product_facebook: "Facebook",
            product_facebook_msg: "Url of Facebook page",
            product_linkedin: "LinkedIn",
            product_linkedin_msg: "Url of LinkedIn page",
            product_slack: "Slack",
            product_slack_msg: "Url of Slack channel",
            product_twitter: "Twitter",
            product_twitter_msg: "Url of Twitter page",

            error_product_not_found: 'Product "{id}" not found.',

            add_product: "Add new product",
            product_added_msg: 'Product "{name}" has been created successfully.',

            delete_product: "Delete product",
            delete_product_msg: 'All topics and pages belong to the product will also be deleted! This product currently has {numTopics} topic(s). Are you sure you wish to delete the product?',
            product_deleted_msg: 'Product "{name}" has been deleted successfully.',

            edit_product: "Edit product",
            product_updated_msg: 'Product "{name}" has been updated successfully.',

            product_unmap_domain: "Unmap",
            product_unmap_domain_msg: 'Are you sure you wish to unmap domain name "{domain}"? Product documents are no longer accessible via this domain name once unmapped.',
            product_domain_unmapped_msg: 'Domain "{domain}" has been unmapped successfully.',
            product_map_domain: "Map",
            product_map_domain_msg: "Product documents are accessible via mapped domain names",
            product_domain_mapped_msg: 'Domain "{domain}" has been mapped successfully.',

            error_topic_not_found: 'Topic "{id}" not found.',

            add_topic: "Add new topic",
            topic_added_msg: 'Topic "{name}" has been created successfully.',

            delete_topic: "Delete topic",
            delete_topic_msg: 'All pages belong to the topic will also be deleted! This topic currently has {numPages} page(s). Are you sure you wish to delete the topic?',
            topic_deleted_msg: 'Topic "{name}" has been deleted successfully.',

            edit_topic: "Edit topic",
            topic_updated_msg: 'Topic "{name}" has been updated successfully.',

            topic_id: "Id",
            topic_id_msg: "Must be unique, will be generated if empty",
            topic_icon: "Icon",
            topic_icon_msg: "Pick an icon from the list",
            topic_title: "Title",
            topic_title_msg: "Topic title",
            topic_summary: "Summary",
            topic_summary_msg: "Short summary of the topic",

            add_page: "Add new page",
            page_added_msg: 'Page "{name}" has been created successfully.',

            delete_page: "Delete page",
            delete_page_msg: 'Are you sure you wish to delete this page?',
            page_deleted_msg: 'Page "{name}" has been deleted successfully.',

            edit_page: "Edit page",
            page_updated_msg: 'Page "{name}" has been updated successfully.',

            page_id: "Id",
            page_id_msg: "Must be unique, will be generated if empty",
            page_icon: "Icon",
            page_icon_msg: "Pick an icon from the list",
            page_title: "Title",
            page_title_msg: "Page title",
            page_summary: "Summary",
            page_summary_msg: "Short summary of the page",
            page_content: "Content",
            page_content_msg: "Page content (Markdown supported)",

            content_editor: "Editor",
            content_preview: "Preview",

            users: "Users",
            my_profile: "My profile",
            user_is_admin: "Is administrator",
            user_is_admin_msg: "Administrator has permission to create other administrator accounts",
            user_id: "Login id",
            user_id_msg: "User email address as login id, must be unique",
            user_mask_id: "Mask id",
            user_mask_id_msg: "User mask-id to not expose user email address, must be unique",
            user_display_name: "Display name",
            user_display_name_msg: "Name of user for displaying purpose",
            user_password: "Password",
            user_current_password: "Current password",
            user_current_password_msg: "Change password: enter current password and the new one",
            user_new_password: "New password",
            user_new_password_msg: "Enter new password. If empty, user can only login using social network account",
            user_confirmed_password: "Confirmed password",
            user_confirmed_password_msg: "Enter new password again to confirm",
            error_confirmed_password_mismatch: "Password does not match the confirmed one.",

            edit_user_profile: "Edit user profile",
            user_profile_updated_msg: 'User profile "{id}" has been updated successfully.',
            user_password_updated_msg: 'Password of user "{id}" has been updated successfully.',

            add_user: "Add new user",
            user_added_msg: 'User "{id}" has been added successfully.',

            delete_user: "Delete user account",
            delete_user_msg: 'Are you sure you wish to remove user account "{id}"?',
            user_deleted_msg: 'User "{id}" has been deleted successfully.',
        }
    },
    vi: {
        _name: 'Ti???ng Vi???t',
        _flag: 'cif-vn',
        _demo_msg: 'B???n tri???n khai n??y d??nh do m???c ????ch tr???i nghi???m. D??? li???u c?? th??? ???????c xo?? v?? c??i ?????t l???i nguy??n g???c b???t c??? l??c n??o m?? kh??ng c???n b??o tr?????c.' +
            '<br/>????ng nh???p v???i t??i kho???n <strong>admin@local/s3cr3t</strong>, ' +
            'ho???c s??? d???ng <i>t??i kho???n mxh</i> b???ng c??ch nh???n v??o ???????ng d???n "????ng nh???p v???i t??i kho???n mxh" ' +
            '(m??y ch??? s??? kh??ng bi???t v?? c??ng kh??ng l??u tr??? th??ng tin ????ng nh???p t??i kho???n mxh c???a b???n).' +
            '<br/><br/>B???n c?? th??? truy c???p v??o trang frontend b???ng <a href="/doc/" style="color: yellowgreen">???????ng d???n n??y</a>.',

        message: {
            actions: 'H??nh ?????ng',
            action_create: 'T???o',
            action_save: 'L??u',
            action_back: 'Quay l???i',
            action_edit: 'S???a',
            action_delete: 'Xo??',
            action_move_down: 'Chuy???n xu???ng',
            action_move_up: 'Chuy???n l??n',

            error: 'C?? l???i',
            info: 'Th??ng tin',
            settings: 'C??i ?????t',
            account: 'T??i kho???n',
            language: 'Ng??n ng???',

            yes: 'C??',
            no: 'Kh??ng',
            ok: '?????ng ??',
            cancel: 'Hu???',
            close: '????ng',

            error_field_mandatory: 'Tr?????ng d??? li???u kh??ng ???????c b??? tr???ng, vui l??ng nh???p th??ng tin.',

            logout: '????ng xu???t',
            login: '????ng nh???p',
            login_info: '????ng nh???p ????? ti???p t???c',
            login_social: '????ng nh???p v???i t??i kho???n mxh',
            wait: 'Vui l??ng gi??? gi??y l??t...',
            error_parse_login_token: 'C?? l???i khi x??? l?? login-token',
            home: 'Trang nh??',
            dashboard: 'T???ng h???p',

            icons: "Bi???u t?????ng",
            icon_icon: "Bi???u t?????ng",
            icon_name: "T??n",

            products: "S???n ph???m",
            topics: "Ch??? ?????",
            pages: "Trang t??i li???u",

            product_info: "Th??ng tin chung",
            product_id: "Id",
            product_id_msg: "Id ph???i l?? duy nh???t, s??? ???????c t??? ?????ng t???o n???u ????? r???ng",
            product_is_published: "????ng t???i",
            product_is_published_msg: "T??i li???u c???a s???n ph???m ch??? xem ???????c khi tr???ng th??i l?? '????ng t???i'",
            product_name: "T??n",
            product_name_msg: "T??n hi???n th??? c???a s???n ph???m",
            product_desc: "M?? t???",
            product_desc_msg: "M?? t??? ng???n v??? s???n ph???m",
            product_domains: "T??n mi???n",
            product_domains_msg: "T??i li???u c???a s???n ph???m truy c???p ???????c t??? c??c t??n mi???n n??y (m???i t??n mi???n 1 d??ng)",

            product_contacts: "Th??ng tin li??n h???",
            product_website: "Website",
            product_website_msg: "?????a ch??? trang web c???a s???n ph???m",
            product_email: "Email",
            product_email_msg: "?????a ch??? email li??n h???",
            product_github: "Github",
            product_github_msg: "Trang GitHub c???a s???n ph???m",
            product_facebook: "Facebook",
            product_facebook_msg: "Trang Facebook c???a s???n ph???m",
            product_linkedin: "LinkedIn",
            product_linkedin_msg: "Trang LinkedIn c???a s???n ph???m",
            product_slack: "Slack",
            product_slack_msg: "Nh??m Slack chat c???a s???n ph???m",
            product_twitter: "Twitter",
            product_twitter_msg: "Trang Twitter c???a s???n ph???m",

            error_product_not_found: 'Kh??ng t??m th???y s???n ph???m "{id}".',

            add_product: "Th??m s???n ph???m",
            product_added_msg: 'S???n ph???m "{name}" ???? ???????c t???o th??nh c??ng.',

            delete_product: "Xo?? s???n ph???m",
            delete_product_msg: 'Xo?? s???n ph???m s??? xo?? c??c ch??? ????? v?? trang t??i li???u c???a s???n ph???m! S???n ph???m n??y hi???n c?? {numTopics} ch??? ?????. B???n c?? ch???c mu???n xo?? s???n ph???m n??y?',
            product_deleted_msg: 'S???n ph???m "{name}" ???? ???????c xo?? th??nh c??ng.',

            edit_product: "Ch???nh s???a s???n ph???m",
            product_updated_msg: 'S???n ph???m "{name}" ???? ???????c c???p nh???t th??nh c??ng.',

            product_unmap_domain: "B??? k???t n???i",
            product_unmap_domain_msg: 'B???n c?? ch???c b??? k???t n???i t??n mi???n "{domain}"? T??i li???u c???a s???n ph???m s??? kh??ng c??n truy c???p ???????c qua t??n mi???n n??y sau khi b??? k???t n???i.',
            product_domain_unmapped_msg: 'K???t n???i v???i t??n mi???n "{domain}" ???? ???????c b??? th??nh c??ng.',
            product_map_domain: "K???t n???i",
            product_map_domain_msg: "T??i li???u c???a s???n ph???m truy c???p ???????c t??? c??c t??n mi???n sau khi ???????c k???t n???i",
            product_domain_mapped_msg: 'T??n mi???n "{domain}" ???? ???????c k???t n???i th??nh c??ng.',

            error_topic_not_found: 'Kh??ng t??m th???y ch??? ????? "{id}".',

            add_topic: "Th??m ch??? ?????",
            topic_added_msg: 'Ch??? ????? "{name}" ???? ???????c t???o th??nh c??ng.',

            delete_topic: "Xo?? ch??? ?????",
            delete_topic_msg: 'Xo?? ch??? ????? s??? xo?? c??c trang t??i li???u n???m trong ch??? ?????! Ch??? ????? n??y hi???n c?? {numPages} tragn t??i li???u. B???n c?? ch???c mu???n xo?? ch??? ????? n??y?',
            topic_deleted_msg: 'Ch??? ????? "{name}" ???? ???????c xo?? th??nh c??ng.',

            edit_topic: "Ch???nh s???a ch??? ?????",
            topic_updated_msg: 'Ch??? ????? "{name}" ???? ???????c c???p nh???t th??nh c??ng.',

            topic_id: "Id",
            topic_id_msg: "Id ph???i l?? duy nh???t, s??? ???????c t??? ?????ng t???o n???u ????? r???ng",
            topic_icon: "Bi???u t?????ng",
            topic_icon_msg: "Ch???n 1 bi???u t?????ng cho ch??? ????? trong danh s??ch",
            topic_title: "T??n",
            topic_title_msg: "T??n hi???n th??? c???a ch??? ?????",
            topic_summary: "T??m t???t",
            topic_summary_msg: "Ph???n t??m t???t ng???n v??? ch??? ?????",

            add_page: "Th??m trang t??i li???u",
            page_added_msg: 'Trang t??i li???u "{name}" ???? ???????c t???o th??nh c??ng.',

            delete_page: "Xo?? trang t??i li???u",
            delete_page_msg: 'B???n c?? ch???c mu???n xo?? trang t??i li???u n??y?',
            page_deleted_msg: 'Trang t??i li???u "{name}" ???? ???????c xo?? th??nh c??ng.',

            edit_page: "Ch???nh s???a trang t??i li???u",
            page_updated_msg: 'Trang t??i li???u "{name}" ???? ???????c c???p nh???t th??nh c??ng.',

            page_id: "Id",
            page_id_msg: "Id ph???i l?? duy nh???t, s??? ???????c t??? ?????ng t???o n???u ????? r???ng",
            page_icon: "Bi???u t?????ng",
            page_icon_msg: "Ch???n 1 bi???u t?????ng cho trang t??i li???u trong danh s??ch",
            page_title: "T??n",
            page_title_msg: "T??n hi???n th??? c???a trang t??i li???u",
            page_summary: "T??m t???t",
            page_summary_msg: "Ph???n t??m t???t ng???n v??? n???i dung c???a trang t??i li???u",
            page_content: "N???i dung",
            page_content_msg: "N???i dung c???a trang t??i li???u (h??? tr??? Markdown)",

            content_editor: "So???n th???o",
            content_preview: "Xem tr?????c",

            users: "Ng?????i d??ng",
            my_profile: "Th??ng tin c?? nh??n",
            user_is_admin: "Qu???n tr??? vi??n",
            user_is_admin_msg: "Qu???n tr??? vi??n s??? ???????c quy???n t???o th??m t??i kho???n qu???n tr??? vi??n kh??c",
            user_id: "T??n ????ng nh???p",
            user_id_msg: "S??? d???ng ?????a ch??? email l??m T??n ????ng nh???p, ph???i l?? duy nh???t tr??n h??? th???ng",
            user_mask_id: "Mask id",
            user_mask_id_msg: "Mask-id s??? ???????c s??? d???ng ????? tr??nh hi???n th??? ?????a ch??? email, ph???i l?? duy nh???t tr??n h??? th???ng",
            user_display_name: "T??n hi???n th???",
            user_display_name_msg: "T??n c???a ng?????i d??ng",
            user_password: "M???t m??",
            user_current_password: "M???t m?? hi???n t???i",
            user_current_password_msg: "????? ?????i m???t m??: nh???p m???t m?? hi???n t???i v?? m???t m?? m???i",
            user_new_password: "M???t m?? m???i",
            user_new_password_msg: "Nh???p m???t m?? m???i. N???u r???ng, ng?????i d??ng ch??? c?? th??? ????ng nh???p th??ng qua t??i kho???n mxh",
            user_confirmed_password: "X??c nh???n l???i m???t m??",
            user_confirmed_password_msg: "Nh???p m???t m?? m???i l???n n???a ????? x??c nh???n",
            error_confirmed_password_mismatch: "M???t m?? kh??ng kh???p nhau",

            edit_user_profile: "Thay ?????i th??ng tin",
            user_profile_updated_msg: 'Th??ng tin ng?????i d??ng "{id}" ???? ???????c c???p nh???t th??nh c??ng.',
            user_password_updated_msg: 'M???t m?? ng?????i d??ng "{id}" ???? ???????c c???p nh???t th??nh c??ng.',

            add_user: "Th??m t??i kho???n ng?????i d??ng",
            user_added_msg: 'Ng?????i d??ng "{id}" ???? ???????c th??m v??o h??? th???ng.',

            delete_user: "Xo?? t??i kho???n ng?????i d??ng",
            delete_user_msg: 'B???n c?? ch???c mu???n xo?? t??i kho???n "{id}" kh???i h??? th???ng?',
            user_deleted_msg: 'T??i kho???n ng?????i d??ng "{id}" ???? ???????c xo?? kh???i h??? th???ng.',
        }
    }
}

Vue.use(VueI18n)

const i18n = new VueI18n({
    locale: 'en',
    messages: messages
})

export default i18n
