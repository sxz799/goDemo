<template>
  <a-row type="flex">
    <a-col :flex="5"></a-col>
    <a-col :flex="5">
      <div style="padding-inline: 20%;">
        <a-card style="text-align: center">
          <h2 style="color: darkgreen">excel格式检测小工具</h2>

          <a-upload-dragger :progress="progress" name="file"
                            :showUploadList="false"
                            :multiple="false" :action="uploadApi" @change="handleChange">
            <p class="ant-upload-drag-icon">
              <inbox-outlined></inbox-outlined>
            </p>
            <p style="color: #ea073a" class="ant-upload-text">点击此处或拖拽文件到这里进行检测</p>
          </a-upload-dragger>

          <h4 style="color: #0bded5">powered by sxz799</h4>
          <a-divider orientation="center"></a-divider>
          <a-table :columns="columns" :data-source="tableData"></a-table>
        </a-card>


      </div>


    </a-col>
    <a-col :flex="5"></a-col>
  </a-row>
</template>

<script>

import {CopyOutlined, DownloadOutlined, InboxOutlined} from '@ant-design/icons-vue';
import {message} from 'ant-design-vue';
import {defineComponent} from 'vue';

export default defineComponent({
  components: {
    InboxOutlined, DownloadOutlined, CopyOutlined
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
    return {
      uploadApi: "/api/upload",
      progress,
      handleChange: e => {
        if (e.file.status === 'done') {


          if (e.file.response.success) {
            message.success(e.file.response.msg)
          }else {
            message.error(e.file.response.msg)
          }
           this.tableData=[]
           this.tableData=e.file.response.errMsgs


        }
      },
      tableData: [],
      columns: [
        {
          title: "错误提醒",
          dataIndex: "msg",
          key: "msg"
        },
      ],
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
