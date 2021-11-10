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

            error: 'Error',
            info: 'Information',

            yes: 'Yes',
            no: 'No',
            ok: 'Ok',
            cancel: 'Cancel',
            close: 'Close',

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

            icons: "Icons",

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

            error_product_not_found: 'Product "{id}" not found.',

            add_product: "Add new product",
            product_added_msg: 'Product "{name}" has been created successfully.',

            delete_product: "Delete product",
            product_deleted_msg: 'Product "{name}" has been deleted successfully.',

            edit_product: "Edit product",
            product_updated_msg: 'Product "{name}" has been updated successfully.',
            product_unmap_domain: "Unmap",
            product_unmap_domain_msg: 'Are you sure you wish to unmap domain name "{domain}"? Product documents are no longer accessible via domain name once unmapped.',
            product_domain_unmapped_msg: 'Domain "{domain}" has been unmapped successfully.',
            product_map_domain: "Map",
            product_map_domain_msg: "Product documents are accessible via mapped domain names",
            product_domain_mapped_msg: 'Domain "{domain}" has been mapped successfully.',

            add_topic: "Add new topic",
            topic_added_msg: 'Topic "{name}" has been created successfully.',

            delete_topic: "Delete topic",
            delete_topic_msg: 'All pages belong to the topic will also be deleted! This topic currently has {numPages} page(s). Are you sure you wish to delete the topic?',
            topic_deleted_msg: 'Topic "{name}" has been deleted successfully.',

            topic_id: "Id",
            topic_id_msg: "Must be unique, will be generated if empty",
            topic_icon: "Icon",
            topic_icon_msg: "Pick an icon from the list",
            topic_title: "Title",
            topic_title_msg: "Topic title",
            topic_summary: "Summary",
            topic_summary_msg: "Short summary of the topic",
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

            error: 'Có lỗi',
            info: 'Thông tin',

            yes: 'Có',
            no: 'Không',
            ok: 'Đồng ý',
            cancel: 'Huỷ',
            close: 'Đóng',

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

            icons: "Icon",

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

            error_product_not_found: 'Không tìm thấy sản phẩm "{id}".',

            add_product: "Thêm sản phẩm",
            product_added_msg: 'Sản phẩm "{name}" đã được tạo thành công.',

            delete_product: "Xoá sản phẩm",
            product_deleted_msg: 'Sản phẩm "{name}" đã được xoá thành công.',

            edit_product: "Chỉnh sửa sản phẩm",
            product_updated_msg: 'Sản phẩm "{name}" đã được cập nhật thành công.',
            product_unmap_domain: "Bỏ kết nối",
            product_unmap_domain_msg: 'Bạn có chắc bỏ kết nối tên miền "{domain}"? Tài liệu của sản phẩm sẽ không còn truy cập được qua tên miền sau khi bỏ kết nối.',
            product_domain_unmapped_msg: 'Kết nối với tên miền "{domain}" đã được bỏ thành công.',
            product_map_domain: "Kết nối",
            product_map_domain_msg: "Tài liệu của sản phẩm truy cập được từ các tên miền sau khi được kết nối",
            product_domain_mapped_msg: 'Tên miền "{domain}" đã được kết nối thành công.',

            add_topic: "Thêm chủ đề",
            topic_added_msg: 'Chủ đề "{name}" đã được tạo thành công.',

            delete_topic: "Xoá chủ đề",
            delete_topic_msg: 'Xoá chủ đề sẽ xoá các trang tài liệu nằm trong chủ đề! Chủ đề này hiện có {numPages} tragn tài liệu. Bạn có chắc muốn xoá chủ đề này?',
            topic_deleted_msg: 'Chủ đề "{name}" đã được xoá thành công.',

            topic_id: "Id",
            topic_id_msg: "Id phải là duy nhất, sẽ được tự động tạo nếu để rỗng",
            topic_icon: "Biểu tượng",
            topic_icon_msg: "Chọn 1 biểu tượng cho chủ đề trong danh sách",
            topic_title: "Tên",
            topic_title_msg: "Tên hiển thị của chủ đề",
            topic_summary: "Mô tả",
            topic_summary_msg: "Phần mô tả ngắn về chủ đề",
        }
    }
}

Vue.use(VueI18n)

const i18n = new VueI18n({
    locale: 'en',
    messages: messages
})

export default i18n
