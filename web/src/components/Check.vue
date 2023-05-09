<template>
  <a-row type="flex">
    <a-col :span="2"></a-col>
    <a-col :span="20">
      <div style="padding-inline: 20%;">
        <a-card style="text-align: center">
          <h2 style="color: #0034ff">资产系统导入模板检验工具</h2>

          <a-upload-dragger :progress="progress" name="file" :before-upload="beforeUpload"
                            :showUploadList="false"
                            accept=".xls,.xlsx"
                            :multiple="false" :action="uploadApi" @change="handleChange">
            <p class="ant-upload-drag-icon">
              <inbox-outlined></inbox-outlined>
            </p>
            <p style="color: #ea073a" class="ant-upload-text">点击这里 或 将文件拖拽到这里 进行检测</p>
          </a-upload-dragger>
          <h4 style="color: #bc08ef" @click="showNotice">powered by sxz799(点我查看使用说明)</h4>

          <a-divider orientation="center">{{ this.fileName}}</a-divider>
          <a-table :locale="tableLocale" :columns="columns" :data-source="tableData"></a-table>
        </a-card>
      </div>
    </a-col>
    <a-col :span="2"></a-col>
  </a-row>

  <a-modal
      :width="1000"
      v-model:visible="visible"
      :maskClosable="false"
      :closable="false"
      :keyboard="false"
      title="使用说明"
      okText="知道了"
      cancel-text="不再提醒"
      @ok="handleOk"
      @cancel="handleHidden"
  >
    <Notice/>
  </a-modal>
</template>

<script>

import {CopyOutlined, DownloadOutlined, InboxOutlined} from '@ant-design/icons-vue';
import {message} from 'ant-design-vue';
import {defineComponent, ref} from 'vue';
import Notice from './Notice.vue'

export default defineComponent({
  components: {
    InboxOutlined, DownloadOutlined, CopyOutlined, Notice
  },

  data() {

    const progress = {
      strokeColor: {
        '0%': '#108ee9',
        '100%': '#87d068',
      },
      strokeWidth: 3,
      format: percent => `${parseFloat(percent.toFixed(2))}%`,
      class: 'test',
    };
    const tableLocale = {
      emptyText: '暂无数据'
    };
    const beforeUpload = file => {
      console.log(file)
      this.fileName=file.name
    };
    return {
      beforeUpload,
      fileName: '',
      tableLocale,
      visible: localStorage.getItem("showAgain") != "N",
      uploadApi: "/api/upload",
      progress,
      handleChange: e => {
        if (e.file.status === 'done') {
          if (e.file.response.success) {
            message.success(e.file.response.msg)
          } else {
            message.error(e.file.response.msg)
          }
          this.tableData = []
          this.tableData = e.file.response.errInfos
        }
      },
      tableData: [],
      columns: [
        {
          title: "行号",
          dataIndex: "line",
          key: "line"
        },
        {
          title: "错误信息",
          dataIndex: "errorMsg",
          key: "errorMsg"
        },
        {
          title: "修改提醒",
          dataIndex: "fixMsg",
          key: "fixMsg"
        },
      ],
    }

  },
  methods: {
    handleOk() {
      this.visible = false;
    },
    handleHidden() {
      this.visible = false;
      localStorage.setItem("showAgain", "N")
    },
    showNotice() {
      this.visible = true;
    }
  }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">
h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
