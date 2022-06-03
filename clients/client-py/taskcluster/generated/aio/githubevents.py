# coding=utf-8
#####################################################
# THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT #
#####################################################
# noqa: E128,E201
from ...aio.asyncclient import AsyncBaseClient
from ...aio.asyncclient import createApiClient
from ...aio.asyncclient import config
from ...aio.asyncclient import createTemporaryCredentials
from ...aio.asyncclient import createSession
_defaultConfig = config


class GithubEvents(AsyncBaseClient):
    """
    The github service publishes a pulse
    message for supported github events, translating Github webhook
    events into pulse messages.

    This document describes the exchange offered by the taskcluster
    github service
    """

    classOptions = {
        "exchangePrefix": "exchange/taskcluster-github/v1/",
    }
    serviceName = 'github'
    apiVersion = 'v1'

    def pullRequest(self, *args, **kwargs):
        """
        GitHub Pull Request Event

        When a GitHub pull request event is posted it will be broadcast on this
        exchange with the designated `organization` and `repository`
        in the routing-key along with event specific metadata in the payload.

        This exchange takes the following keys:

         * routingKeyKind: Identifier for the routing-key kind. This is always `"primary"` for the formalized routing key. (required)

         * organization: The GitHub `organization` which had an event. All periods have been replaced by % - such that foo.bar becomes foo%bar - and all other special characters aside from - and _ have been stripped. (required)

         * repository: The GitHub `repository` which had an event.All periods have been replaced by % - such that foo.bar becomes foo%bar - and all other special characters aside from - and _ have been stripped. (required)

         * action: The GitHub `action` which triggered an event. See for possible values see the payload actions property. (required)
        """

        ref = {
            'exchange': 'pull-request',
            'name': 'pullRequest',
            'routingKey': [
                {
                    'constant': 'primary',
                    'multipleWords': False,
                    'name': 'routingKeyKind',
                },
                {
                    'multipleWords': False,
                    'name': 'organization',
                },
                {
                    'multipleWords': False,
                    'name': 'repository',
                },
                {
                    'multipleWords': False,
                    'name': 'action',
                },
            ],
            'schema': 'v1/github-pull-request-message.json#',
        }
        return self._makeTopicExchange(ref, *args, **kwargs)

    def push(self, *args, **kwargs):
        """
        GitHub push Event

        When a GitHub push event is posted it will be broadcast on this
        exchange with the designated `organization` and `repository`
        in the routing-key along with event specific metadata in the payload.

        This exchange takes the following keys:

         * routingKeyKind: Identifier for the routing-key kind. This is always `"primary"` for the formalized routing key. (required)

         * organization: The GitHub `organization` which had an event. All periods have been replaced by % - such that foo.bar becomes foo%bar - and all other special characters aside from - and _ have been stripped. (required)

         * repository: The GitHub `repository` which had an event.All periods have been replaced by % - such that foo.bar becomes foo%bar - and all other special characters aside from - and _ have been stripped. (required)
        """

        ref = {
            'exchange': 'push',
            'name': 'push',
            'routingKey': [
                {
                    'constant': 'primary',
                    'multipleWords': False,
                    'name': 'routingKeyKind',
                },
                {
                    'multipleWords': False,
                    'name': 'organization',
                },
                {
                    'multipleWords': False,
                    'name': 'repository',
                },
            ],
            'schema': 'v1/github-push-message.json#',
        }
        return self._makeTopicExchange(ref, *args, **kwargs)

    def release(self, *args, **kwargs):
        """
        GitHub release Event

        When a GitHub release event is posted it will be broadcast on this
        exchange with the designated `organization` and `repository`
        in the routing-key along with event specific metadata in the payload.

        This exchange takes the following keys:

         * routingKeyKind: Identifier for the routing-key kind. This is always `"primary"` for the formalized routing key. (required)

         * organization: The GitHub `organization` which had an event. All periods have been replaced by % - such that foo.bar becomes foo%bar - and all other special characters aside from - and _ have been stripped. (required)

         * repository: The GitHub `repository` which had an event.All periods have been replaced by % - such that foo.bar becomes foo%bar - and all other special characters aside from - and _ have been stripped. (required)
        """

        ref = {
            'exchange': 'release',
            'name': 'release',
            'routingKey': [
                {
                    'constant': 'primary',
                    'multipleWords': False,
                    'name': 'routingKeyKind',
                },
                {
                    'multipleWords': False,
                    'name': 'organization',
                },
                {
                    'multipleWords': False,
                    'name': 'repository',
                },
            ],
            'schema': 'v1/github-release-message.json#',
        }
        return self._makeTopicExchange(ref, *args, **kwargs)

    def rerun(self, *args, **kwargs):
        """
        GitHub re-run task Event

        When a GitHub check_run event with action="rerequested" is posted
        it will be broadcast on this exchange with the designated
        `organization` and `repository`
        in the routing-key along with event specific metadata in the payload.

        This exchange takes the following keys:

         * routingKeyKind: Identifier for the routing-key kind. This is always `"primary"` for the formalized routing key. (required)

         * organization: The GitHub `organization` which had an event. All periods have been replaced by % - such that foo.bar becomes foo%bar - and all other special characters aside from - and _ have been stripped. (required)

         * repository: The GitHub `repository` which had an event.All periods have been replaced by % - such that foo.bar becomes foo%bar - and all other special characters aside from - and _ have been stripped. (required)
        """

        ref = {
            'exchange': 'rerun',
            'name': 'rerun',
            'routingKey': [
                {
                    'constant': 'primary',
                    'multipleWords': False,
                    'name': 'routingKeyKind',
                },
                {
                    'multipleWords': False,
                    'name': 'organization',
                },
                {
                    'multipleWords': False,
                    'name': 'repository',
                },
            ],
            'schema': 'v1/github-rerun-message.json#',
        }
        return self._makeTopicExchange(ref, *args, **kwargs)

    def taskGroupCreationRequested(self, *args, **kwargs):
        """
        tc-gh requested the Queue service to create all the tasks in a group

        supposed to signal that taskCreate API has been called for every task in the task group
        for this particular repo and this particular organization
        currently used for creating initial status indicators in GitHub UI using Statuses API.
        This particular exchange can also be bound to RabbitMQ queues by custom routes - for that,
        Pass in the array of routes as a second argument to the publish method. Currently, we do
        use the statuses routes to bind the handler that creates the initial status.

        This exchange takes the following keys:

         * routingKeyKind: Identifier for the routing-key kind. This is always `"primary"` for the formalized routing key. (required)

         * organization: The GitHub `organization` which had an event. All periods have been replaced by % - such that foo.bar becomes foo%bar - and all other special characters aside from - and _ have been stripped. (required)

         * repository: The GitHub `repository` which had an event.All periods have been replaced by % - such that foo.bar becomes foo%bar - and all other special characters aside from - and _ have been stripped. (required)
        """

        ref = {
            'exchange': 'task-group-creation-requested',
            'name': 'taskGroupCreationRequested',
            'routingKey': [
                {
                    'constant': 'primary',
                    'multipleWords': False,
                    'name': 'routingKeyKind',
                },
                {
                    'multipleWords': False,
                    'name': 'organization',
                },
                {
                    'multipleWords': False,
                    'name': 'repository',
                },
            ],
            'schema': 'v1/task-group-creation-requested.json#',
        }
        return self._makeTopicExchange(ref, *args, **kwargs)

    funcinfo = {
    }


__all__ = ['createTemporaryCredentials', 'config', '_defaultConfig', 'createApiClient', 'createSession', 'GithubEvents']
