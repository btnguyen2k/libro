//#GovueAdmin-Customized
import Vue from 'vue'
import VueI18n from 'vue-i18n'

const messages = {
    en: {
        message: {
            actions: 'Actions',
            action_create: 'Create',
            action_save: 'Save',
            action_back: 'Back',
            action_edit: 'Edit',
            action_delete: 'Delete',

            login: 'Login',
            login_info: 'Please sign in to continue',
            login_social: 'Login with social account',
            username: 'Username',
            password: 'Password',
            demo_msg: "This is instance is for demo purpose only. Login with default account <strong>admin@local/s3cr3t</strong>.<br/>Or you can login with your <u>social account</u> via \"Login with social account\" link (your social account credential <u>will not</u> be stored on the server).",
            wait: 'Please wait...',
            error_parse_login_token: 'Error parsing login-token',
            home: 'Home',
            dashboard: 'Dashboard',

            products: "Products",
            topics: "Topics",
            pages: "Pages",

            product_is_published: "Published",
            product_is_published_msg: "Product's documents are visible only when published",
            product_name: "Name",
            product_name_msg: "Display name of the product",
            product_desc: "Description",
            product_desc_msg: "Summary description of the product",
            product_domains: "Domain names",
            product_domains_msg: "Product documents are accessible via these domain names (one domain per line)",

            add_product: "Add new product",
            product_added_msg: 'Product "{name}" has been created successfully.'
        }
    },
    vi: {
        message: {
            actions: 'Hành động',
            action_create: 'Tạo',
            action_save: 'Lưu',
            action_back: 'Quay lại',
            action_edit: 'Sửa',
            action_delete: 'Xoá',

            login: 'Đăng nhập',
            login_info: 'Đăng nhập để tiếp tục',
            login_social: 'Đăng nhập với tài khoản mxh',
            username: 'Tên đăng nhập',
            password: 'Mật mã',
            demo_msg: 'Bản triển khai này dành do mục đích thử nghiệm. Đăng nhập với tài khoản <strong>admin@local/s3cr3t</strong>.<br/>Hoặc đăng nhập với <i>tài khoản mxh</i> (nhấn vào đường dẫn "Đăng nhập với tài khoản mxh").',
            wait: 'Vui lòng giờ giây lát...',
            error_parse_login_token: 'Có lỗi khi xử lý login-token',
            home: 'Trang nhà',
            dashboard: 'Tổng hợp',

            products: "Sản phẩm",
            topics: "Chủ đề",
            pages: "Trang tài liệu",

            product_is_published: "Đăng tải",
            product_is_published_msg: "Tài liệu của sản phẩm chỉ xem được khi trạng thái là 'Đăng tải'",
            product_name: "Tên",
            product_name_msg: "Tên hiển thị của sản phẩm",
            product_desc: "Mô tả",
            product_desc_msg: "Mô tả ngắn về sản phẩm",
            product_domains: "Tên miền",
            product_domains_msg: "Tài liệu của sản phẩm truy cập được từ các tên miền này (mỗi tên miền 1 dòng)",

            add_product: "Thêm sản phẩm",
            product_added_msg: 'Sản phẩm "{name}" đã được tạo thành công.'
        }
    }
}

Vue.use(VueI18n)

const i18n = new VueI18n({
    locale: 'en',
    messages: messages
})

export default i18n