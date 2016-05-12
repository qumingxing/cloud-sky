package main

import (
	/*
		#include <libavcodec/avcodec.h>
		#include <libavutil/avutil.h>
		#include <string.h>
		#include <stdio.h>
		typedef struct {
			AVCodec *c;
			AVCodecContext *ctx;
			AVFrame *f;
			int got;
		} aacdec_t ;
		static int aacdec_new(aacdec_t *m, uint8_t *buf, int len) {
			m->c = avcodec_find_decoder(AV_CODEC_ID_AAC);
			m->ctx = avcodec_alloc_context3(m->c);
			m->f = av_frame_alloc();
			m->ctx->extradata = buf;
			m->ctx->extradata_size = len;
			m->ctx->debug = 0x3;
			av_log(m->ctx, AV_LOG_DEBUG, "m %p\n", m);
			return avcodec_open2(m->ctx, m->c, 0);
		}
		static int aacdec_decode(aacdec_t *m, uint8_t *data, int len) {
			AVPacket pkt;
			av_init_packet(&pkt);
			pkt.data = data;
			pkt.size = len;
			av_log(m->ctx, AV_LOG_DEBUG, "decode %p\n", m);
			return avcodec_decode_audio4(m->ctx, m->f, &m->got, &pkt);
		}
	*/
	"C"
	"fmt"
)
type AACDecoder struct {
	m C.aacdec_t
}
var (
	aav string

)
func main(){
	var aaa string="adsfads"
	var bbb string="adf"
	var dd =""
	fmt.Print(aaa,bbb,dd)
}
