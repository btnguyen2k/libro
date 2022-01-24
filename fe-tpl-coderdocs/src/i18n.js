//#Libro frontend, template CoderDocs
import Vue from 'vue'
import VueI18n from 'vue-i18n'

const messages = {
    en: {
        wait: 'Please wait...',
        click_to_home: 'Click here to navigate to the home page.',
        click_to_topic: 'Click here to navigate to the topic index.',
        contacts: 'Contacts',
        topics: 'Topics',
        pages: 'Pages',

        page_updated_timestamp: 'Last updated',
        page_created_timestamp: 'Creation time',

        error_unknown: "Error occurred, try again after a moment! If the problem persists, please contact website administrator.",
        error_product_not_found: 'No product mapped to domain "{domain}".',
        error_topic_not_found: 'Topic "{topic}" not found.',
        error_page_not_found: 'Page "{page}" not found.',

        empty_topic: 'This topic has no document page.',
    },
    vi: {
        wait: 'Vui lòng giờ giây lát...',
        click_to_home: 'Nhấn vào đây để đi đến trang chủ.',
        click_to_topic: 'Nhấn vào đây để đi đến trang nhà của chủ đề.',
        contacts: 'Liên hệ',
        topics: 'Chủ đề',
        pages: 'Trang tài liệu',

        page_updated_timestamp: 'Cập nhật lần cuối',
        page_created_timestamp: 'Thời điểm tạo',

        error_unknown: "Có lỗi xảy ra, vui lòng thử lại sau. Nếu sự cố vẫn còn tiếp diễn, hãy liên lạc với người quản trị website.",
        error_product_not_found: 'Không có sản phẩm nào tương ứng với tên miền "{domain}".',
        error_topic_not_found: 'Không tìm thấy chủ đề "{topic}".',
        error_page_not_found: 'Không tìm thấy trang tài liệu "{page}".',

        empty_topic: 'Chủ đề này hiện chưa có bài viết nào.',
    }
}

Vue.use(VueI18n)

const i18n = new VueI18n({
    locale: 'en',
    messages: messages
})

export default i18n
