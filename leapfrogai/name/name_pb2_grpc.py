# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from . import name_pb2 as name_dot_name__pb2


class NameServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Name = channel.unary_unary(
            "/name.NameService/Name",
            request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            response_deserializer=name_dot_name__pb2.NameResponse.FromString,
        )


class NameServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Name(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details("Method not implemented!")
        raise NotImplementedError("Method not implemented!")


def add_NameServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
        "Name": grpc.unary_unary_rpc_method_handler(
            servicer.Name,
            request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            response_serializer=name_dot_name__pb2.NameResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        "name.NameService", rpc_method_handlers
    )
    server.add_generic_rpc_handlers((generic_handler,))


# This class is part of an EXPERIMENTAL API.
class NameService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Name(
        request,
        target,
        options=(),
        channel_credentials=None,
        call_credentials=None,
        insecure=False,
        compression=None,
        wait_for_ready=None,
        timeout=None,
        metadata=None,
    ):
        return grpc.experimental.unary_unary(
            request,
            target,
            "/name.NameService/Name",
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            name_dot_name__pb2.NameResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
        )
