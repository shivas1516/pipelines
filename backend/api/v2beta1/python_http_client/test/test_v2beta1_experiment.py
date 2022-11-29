# coding: utf-8

"""
    Kubeflow Pipelines API

    This file contains REST API specification for Kubeflow Pipelines. The file is autogenerated from the swagger definition.

    Contact: kubeflow-pipelines@google.com
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import kfp_server_api
from kfp_server_api.models.v2beta1_experiment import V2beta1Experiment  # noqa: E501
from kfp_server_api.rest import ApiException

class TestV2beta1Experiment(unittest.TestCase):
    """V2beta1Experiment unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V2beta1Experiment
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = kfp_server_api.models.v2beta1_experiment.V2beta1Experiment()  # noqa: E501
        if include_optional :
            return V2beta1Experiment(
                experiment_id = '0', 
                display_name = '0', 
                description = '0', 
                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                namespace = '0', 
                storage_state = 'STORAGESTATE_UNSPECIFIED'
            )
        else :
            return V2beta1Experiment(
        )

    def testV2beta1Experiment(self):
        """Test V2beta1Experiment"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
