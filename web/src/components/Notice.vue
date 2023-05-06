<template>
  <h2 style="color: #3260c5">为了节省导入时间和提高效率,在此处进行excel表格的初步格式检测</h2>
  <a-collapse>
    <a-collapse-panel key="1" header="第一步 删除多余的sheet页">
      <a-image :width="200" :src="delsheet"/>
    </a-collapse-panel>
    <a-collapse-panel key="2" header="第二步 设置列宽">
      <a-image :width="200" :src="setliekuan"/>

    </a-collapse-panel>
    <a-collapse-panel key="3" header="第三步 设置行高" >
      <a-image :width="200" :src="sethanggao"/>
    </a-collapse-panel>

    <a-collapse-panel key="4" header="第四步 设置单元格格式" >
      <a-image-preview-group>
        <a-image :width="200" :src="setcell1"/>
        <a-image :width="200" :src="setcell2"/>
      </a-image-preview-group>
    </a-collapse-panel>
    <a-collapse-panel key="5" header="第五步 合并前两行" >
      <a-image :width="200" :src="settop2row"/>
    </a-collapse-panel>
    <a-collapse-panel key="6" header="第六步 导出xlsx文件" >
      <a-image :width="200" :src="exportxlsx"/>
    </a-collapse-panel>
    <a-collapse-panel key="7" header="第七步 将导出的文件拖进窗口进行检测" >
    </a-collapse-panel>
  </a-collapse>


</template>

<script>

import delsheet from '@/images/delsheet.png'
import sethanggao from '@/images/sethanggao.png'
import setliekuan from '@/images/setliekuan.png'
import setcell1 from '@/images/setcell1.png'
import setcell2 from '@/images/setcell2.png'
import settop2row from '@/images/settop2row.png'
import exportxlsx from '@/images/exportxlsx.png'

export default {
  name: "Notice",
  data() {
    return {
      delsheet: delsheet,
      sethanggao: sethanggao,
      setliekuan: setliekuan,
      setcell1: setcell1,
      setcell2: setcell2,
      settop2row: settop2row,
      exportxlsx: exportxlsx,
    }

  }
}
</script>

<style scoped>

</style>
